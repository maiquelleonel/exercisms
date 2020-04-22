from datetime import *
from dateutil.relativedelta import *

WEEKDAYS = {
    "Monday": MO,
    "Tuesday": TU,
    "Wednesday": WE,
    "Thursday": TH,
    "Friday": FR,
    "Saturday": SA,
    "Sunday": SU,
}

ORDS = {"last": -1, "1st": 1, "2nd": 2, "3rd": 3, "4th": 4, "5th": 5}


class MeetupDayException(Exception):
    ...


def meetup(year, month, week, day_of_week):
    _date = date(year, month, 1)
    if week == "teenth":
        _date += relativedelta(day=13, weekday=WEEKDAYS[day_of_week])
    elif week == "last":
        _date += relativedelta(day=31, weekday=WEEKDAYS[day_of_week](ORDS[week]))
    else:
        _date += relativedelta(weekday=WEEKDAYS[day_of_week](ORDS[week]))

    if _date.month > month:
        raise MeetupDayException("Month doesn't have this week day")

    return _date
