def convert(number):
    factors = {3: "Pling", 5: "Plang", 7: "Plong"}
    sounds = ""
    for divisor, sound in factors.items():
        if number % divisor == 0:
            sounds += sound
    return sounds if len(sounds) > 0 else str(number)
