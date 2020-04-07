import re


class Luhn:
    regex = r"[^0-9\s]"

    def __init__(self, card_number):
        self.card_number = card_number.strip()

    def _clean_card(self):
        return re.sub(self.regex, "", self.card_number).replace(" ", "")

    def _doesnt_have_valid_lenth(self):
        return len(self.card_number) <= 1

    def _doesnt_have_valid_chars(self):
        return len(re.findall(self.regex, self.card_number)) > 0

    def _calculate(self):
        card_number = self._clean_card()
        positions = [int(i) for i in range(len(card_number) - 2, -1, -2)]
        if len(positions) == 0:
            positions.append(0)

        for_calculate = [int(char) for char in card_number]
        for position in positions:
            product = for_calculate[position] * 2
            for_calculate[position] = product if product < 9 else product - 9

        return sum(for_calculate) % 10 == 0

    def valid(self):
        if self._doesnt_have_valid_lenth() or self._doesnt_have_valid_chars():
            return False
        return self._calculate()
