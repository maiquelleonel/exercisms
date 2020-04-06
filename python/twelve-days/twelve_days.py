def recite(start_verse, end_verse):
    verses = [
        "a Partridge in a Pear Tree.",
        "two Turtle Doves, ",
        "three French Hens, ",
        "four Calling Birds, ",
        "five Gold Rings, ",
        "six Geese-a-Laying, ",
        "seven Swans-a-Swimming, ",
        "eight Maids-a-Milking, ",
        "nine Ladies Dancing, ",
        "ten Lords-a-Leaping, ",
        "eleven Pipers Piping, ",
        "twelve Drummers Drumming, ",
    ]
    keys = [
        "first",
        "second",
        "third",
        "fourth",
        "fifth",
        "sixth",
        "seventh",
        "eighth",
        "ninth",
        "tenth",
        "eleventh",
        "twelfth",
    ]

    template = "On the {number} day of Christmas my true love gave to me: {list}"

    def _format_verses(list):
        if len(list) > 1:
            list[-1] = "and {}".format(list[-1])
        return "".join(list)

    return [
        template.format(number=keys[i - 1], list=_format_verses(verses[:i][::-1]))
        for i in range(start_verse, end_verse + 1)
    ]
