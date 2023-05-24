package memorystorage

import (
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	evt := storage.Event{
		Title:    "Title",
		Time:     time.Now(),
		Duration: time.Hour,
		UserID:   uuid.New(),
	}

	t.Run("create", func(t *testing.T) {
		memStorage := New()
		evtID, _ := memStorage.Create(&evt)
		require.NotEqual(t, uuid.Nil, evtID)
	})

	t.Run("update", func(t *testing.T) {
		memStorage := New()
		err := memStorage.Update(uuid.New(), &storage.Event{})
		require.ErrorIs(t, storage.ErrEventNotFound, err)

		evtID, _ := memStorage.Create(&evt)
		err = memStorage.Update(evtID, &storage.Event{Title: "NewTitle"})
		require.NoError(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		memStorage := New()
		err := memStorage.Delete(uuid.New())
		require.ErrorIs(t, storage.ErrEventNotFound, err)

		evtID, _ := memStorage.Create(&evt)
		err = memStorage.Delete(evtID)
		require.NoError(t, err)
	})

	/* 1970, Jan
	   Mo Tu We Th Fr Sa Su
	   	        1  2  3  4
	   5  6  7  8  9  10 11
	   12 13 14 15 16 17 18
	   19 20 21 22 23 24 25
	   26 27 28 29 30 31
	*/
	t.Run("read", func(t *testing.T) {
		memStorage := New()

		tm := time.Date(1970, time.January, 10, 10, 0, 0, 0, time.UTC)
		firstEvID, _ := memStorage.Create(&storage.Event{
			Time: tm, Title: "0",
		})
		require.NotEqual(t, uuid.Nil, firstEvID)

		secondEvID, _ := memStorage.Create(&storage.Event{
			Time: tm, Title: "1",
		})
		require.NotEqual(t, uuid.Nil, secondEvID)

		de, _ := memStorage.DayEvens(tm)
		require.Equal(t, 2, len(de))
		sort.Slice(de, func(i, j int) bool {
			return de[i].Title < de[j].Title
		})

		require.Equal(t, firstEvID, de[0].ID)
		require.Equal(t, secondEvID, de[1].ID)
		require.Equal(t, "0", de[0].Title)
		require.Equal(t, "1", de[1].Title)

		we, _ := memStorage.WeekEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
		sort.Slice(we, func(i, j int) bool {
			return we[i].Title < we[j].Title
		})
		require.Equal(t, de, we)

		me, _ := memStorage.MonthEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
		sort.Slice(we, func(i, j int) bool {
			return me[i].Title < me[j].Title
		})
		require.Equal(t, we, me)

		memStorage.Delete(secondEvID)
		de, _ = memStorage.DayEvens(tm)
		require.Equal(t, 1, len(de))
		require.Equal(t, firstEvID, de[0].ID)

		memStorage.Delete(firstEvID)
		de, _ = memStorage.DayEvens(tm)
		require.Equal(t, 0, len(de))
	})

	t.Run("complex", func(t *testing.T) {
		memStorage := New()
		days := []int{6, 7, 14, 22, 23, 31}
		for _, day := range days {
			memStorage.Create(&storage.Event{
				Time:  time.Date(1970, time.January, day, 10, 0, 0, 0, time.UTC),
				Title: strconv.Itoa(day),
			})
		}
		me, _ := memStorage.MonthEvens(time.Date(1970, time.January, 1, 10, 0, 0, 0, time.UTC))
		require.Equal(t, len(days), len(me))
		sort.Slice(me, func(i, j int) bool {
			return me[i].Time.Before(me[j].Time)
		})
		for i, e := range me {
			require.Equal(t, days[i], e.Time.Day())
			require.Equal(t, strconv.Itoa(days[i]), e.Title)
		}

		we, _ := memStorage.WeekEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
		require.Equal(t, 2, len(we))
		sort.Slice(me, func(i, j int) bool {
			return me[i].Time.Before(me[j].Time)
		})
		weekDays := []int{6, 7}
		for i, e := range we {
			require.Equal(t, weekDays[i], e.Time.Day())
			require.Equal(t, strconv.Itoa(weekDays[i]), e.Title)
		}
	})
}
