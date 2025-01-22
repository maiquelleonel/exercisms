import datetime

class Clock:
    time_in_seconds = 0

    def __init__(self, hour, minute):
        self.time_in_seconds = (hour * 3600) + (minute * 60)

    def __repr__(self):
        return datetime.datetime.fromtimestamp(self.time_in_seconds, datetime.UTC).strftime("%H:%M")

    def __eq__(self, other):
        return self.__class__ == other.__class__ and str(self) == str(other)

    def __add__(self, minutes):
        self.time_in_seconds += minutes * 60
        return self

    def __sub__(self, minutes):
        self.time_in_seconds -= minutes * 60
        return self
