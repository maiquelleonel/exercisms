import datetime
import calendar


WEEKDAYS = {
    "Monday": calendar.MONDAY,
    "Tuesday": calendar.TUESDAY,
    "Wednesday": calendar.WEDNESDAY,
    "Thursday": calendar.THURSDAY,
    "Friday": calendar.FRIDAY,
    "Saturday": calendar.SATURDAY,
    "Sunday": calendar.SUNDAY,
}

ORDS = {"1st": 0, "2nd": 7, "3rd": 14, "4th": 21, "5th": 28, "last": None}


class MeetupDayException(Exception):
    ...



def meetup(year, month, week, day_of_week):
    _, num_days = calendar.monthrange(year, month)
    if week in list(ORDS.keys()):
        for day in range(1, 8):
            weekday = calendar.weekday(year, month, day)
            if weekday == WEEKDAYS[day_of_week]:
                if week == "last":
                    is_5th_the_last_week = day + ORDS["5th"] <= num_days
                    return datetime.date(
                        year, 
                        month, 
                        day + ORDS["5th" if is_5th_the_last_week else "4th"]
                    )
                else:
                    if day + ORDS[week] <= num_days:
                        return datetime.date(year, month, day + ORDS[week])
    
    elif week == "teenth":
        for day in range(13, 20):
            weekday = calendar.weekday(year, month, day)
            if weekday == WEEKDAYS[day_of_week]:
                return datetime.date(year, month, day)
     
    raise MeetupDayException("That day does not exist.")
