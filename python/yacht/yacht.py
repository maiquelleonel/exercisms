"""
This exercise stub and the test suite contain several enumerated constants.

Since Python 2 does not have the enum module, the idiomatic way to write
enumerated constants has traditionally been a NAME assigned to an arbitrary,
but unique value. An integer is traditionally used because it’s memory
efficient.
It is a common practice to export both constants and functions that work with
those constants (ex. the constants in the os, subprocess and re modules).

You can learn more here: https://en.wikipedia.org/wiki/Enumerated_type
"""


# Score categories.
# Change the values as you see fit.
(
    YACHT,
    ONES,
    TWOS,
    THREES,
    FOURS,
    FIVES,
    SIXES,
    FULL_HOUSE,
    FOUR_OF_A_KIND,
    LITTLE_STRAIGHT,
    BIG_STRAIGHT,
    CHOICE,
) = range(0, 12)


def score(dice, category):
    methods = {
        0: lambda dice: 50 if total_sides(dice).get(5, 0) > 0 else 0,
        7: lambda dice: sum(dice) if is_full_house(dice) else 0,
        8: lambda dice: max(dice) * 4 if is_four_a_kind(dice) else 0,
        9: lambda dice: 30 if sorted(dice) == [1, 2, 3, 4, 5] else 0,
        10: lambda dice: 30 if sorted(dice) == [2, 3, 4, 5, 6] else 0,
        11: lambda dice: sum(dice),
    }

    def calc_score(dice, side, score):
        return dice.count(side) * score

    def total_sides(dice):
        return {dice.count(side): side for side in range(1, 7)}

    def is_full_house(dice):
        sides = total_sides(dice)
        return sides.get(2, 0) > 0 and sides.get(3, 0) > 0

    def is_four_a_kind(dice):
        sides = total_sides(dice)
        return sides.get(4, 0) > 0 or sides.get(5, 0) > 0

    return (
        calc_score(dice, side=category, score=category)
        if category in range(1, 7)
        else methods[category](dice)
    )
