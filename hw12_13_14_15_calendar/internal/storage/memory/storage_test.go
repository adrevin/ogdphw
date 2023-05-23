package memorystorage

import (
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	evt := entities.Event{
		Title:    "Title",
		Time:     time.Now().UTC(),
		Duration: time.Hour,
		UserID:   uuid.New(),
	}

	t.Run("create", func(t *testing.T) {
		storage := New()
		evtID := storage.Create(&evt)
		require.NotEqual(t, uuid.Nil, evtID)
	})

	t.Run("update", func(t *testing.T) {
		storage := New()
		err := storage.Update(uuid.New(), &entities.Event{})
		require.ErrorIs(t, ErrEventNotFound, err)

		evtID := storage.Create(&evt)
		err = storage.Update(evtID, &entities.Event{Title: "NewTitle"})
		require.NoError(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		storage := New()
		err := storage.Delete(uuid.New())
		require.ErrorIs(t, ErrEventNotFound, err)

		evtID := storage.Create(&evt)
		err = storage.Delete(evtID)
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
		storage := New()

		tm := time.Date(1970, time.January, 10, 10, 0, 0, 0, time.UTC)
		firstEvUD := storage.Create(&entities.Event{
			Time: tm, Title: "1970-01-10 - 1",
		})
		require.NotEqual(t, uuid.Nil, firstEvUD)

		secondEvUD := storage.Create(&entities.Event{
			Time: tm, Title: "1970-01-10 - 2",
		})
		require.NotEqual(t, uuid.Nil, secondEvUD)

		de := storage.DayEvens(tm)
		require.Equal(t, 2, len(de))
		require.Equal(t, firstEvUD, de[0].ID)
		require.Equal(t, secondEvUD, de[1].ID)
		require.Equal(t, "1970-01-10 - 1", de[0].Title)
		require.Equal(t, "1970-01-10 - 2", de[1].Title)

		we := storage.WeekEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
		require.Equal(t, de, we)

		me := storage.MonthEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
		require.Equal(t, we, me)

		storage.Delete(secondEvUD)
		de = storage.DayEvens(tm)
		require.Equal(t, 1, len(de))
		require.Equal(t, firstEvUD, de[0].ID)

		storage.Delete(firstEvUD)
		de = storage.DayEvens(tm)
		require.Equal(t, 0, len(de))
	})

	t.Run("complex", func(t *testing.T) {
		storage := New()
		days := []int{6, 7, 14, 22, 23, 31}
		for _, day := range days {
			storage.Create(&entities.Event{
				Time:  time.Date(1970, time.January, day, 10, 0, 0, 0, time.UTC),
				Title: strconv.Itoa(day),
			})
		}
		me := storage.MonthEvens(time.Date(1970, time.January, 1, 10, 0, 0, 0, time.UTC))
		require.Equal(t, len(days), len(me))
		sort.Slice(me, func(i, j int) bool {
			return me[i].Time.Before(me[j].Time)
		})
		for i, e := range me {
			require.Equal(t, days[i], e.Time.Day())
			require.Equal(t, strconv.Itoa(days[i]), e.Title)
		}

		we := storage.WeekEvens(time.Date(1970, time.January, 5, 10, 0, 0, 0, time.UTC))
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
