package resistorcolortrio

// This is an auto-generated file. Do not change it manually. Run the generator to update the file.
// See https://github.com/exercism/go#synchronizing-tests-and-instructions
// Source: exercism/problem-specifications
// Commit: cb7eab5 Resistor Color Trio: extend tests, borrowing from awk (#2162)

type labelTestCase struct {
	description string
	input       []string
	expected    string
}

var labelTestCases = []labelTestCase{
	{
		description: "Orange and orange and black",
		input:       []string{"orange", "orange", "black"},
		expected:    "33 ohms",
	},
	{
		description: "Blue and grey and brown",
		input:       []string{"blue", "grey", "brown"},
		expected:    "680 ohms",
	},
	{
		description: "Red and black and red",
		input:       []string{"red", "black", "red"},
		expected:    "2 kiloohms",
	},
	{
		description: "Green and brown and orange",
		input:       []string{"green", "brown", "orange"},
		expected:    "51 kiloohms",
	},
	{
		description: "Yellow and violet and yellow",
		input:       []string{"yellow", "violet", "yellow"},
		expected:    "470 kiloohms",
	},
	{
		description: "Blue and violet and blue",
		input:       []string{"blue", "violet", "blue"},
		expected:    "67 megaohms",
	},
	{
		description: "Minimum possible value",
		input:       []string{"black", "black", "black"},
		expected:    "0 ohms",
	},
	{
		description: "Maximum possible value",
		input:       []string{"white", "white", "white"},
		expected:    "99 gigaohms",
	},
	{
		description: "First two colors make an invalid octal number",
		input:       []string{"black", "grey", "black"},
		expected:    "8 ohms",
	},
	{
		description: "Ignore extra colors",
		input:       []string{"blue", "green", "yellow", "orange"},
		expected:    "650 kiloohms",
	},
	{
		description: "black and black and black",
		input:       []string{"black", "black", "black"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and brown",
		input:       []string{"black", "black", "brown"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and red",
		input:       []string{"black", "black", "red"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and orange",
		input:       []string{"black", "black", "orange"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and yellow",
		input:       []string{"black", "black", "yellow"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and green",
		input:       []string{"black", "black", "green"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and blue",
		input:       []string{"black", "black", "blue"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and violet",
		input:       []string{"black", "black", "violet"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and grey",
		input:       []string{"black", "black", "grey"},
		expected:    "0 ohms",
	},
	{
		description: "black and black and white",
		input:       []string{"black", "black", "white"},
		expected:    "0 ohms",
	},
	{
		description: "black and brown and black",
		input:       []string{"black", "brown", "black"},
		expected:    "1 ohms",
	},
	{
		description: "black and brown and brown",
		input:       []string{"black", "brown", "brown"},
		expected:    "10 ohms",
	},
	{
		description: "black and brown and red",
		input:       []string{"black", "brown", "red"},
		expected:    "100 ohms",
	},
	{
		description: "black and brown and orange",
		input:       []string{"black", "brown", "orange"},
		expected:    "1 kiloohms",
	},
	{
		description: "black and brown and yellow",
		input:       []string{"black", "brown", "yellow"},
		expected:    "10 kiloohms",
	},
	{
		description: "black and brown and green",
		input:       []string{"black", "brown", "green"},
		expected:    "100 kiloohms",
	},
	{
		description: "black and brown and blue",
		input:       []string{"black", "brown", "blue"},
		expected:    "1 megaohms",
	},
	{
		description: "black and brown and violet",
		input:       []string{"black", "brown", "violet"},
		expected:    "10 megaohms",
	},
	{
		description: "black and brown and grey",
		input:       []string{"black", "brown", "grey"},
		expected:    "100 megaohms",
	},
	{
		description: "black and brown and white",
		input:       []string{"black", "brown", "white"},
		expected:    "1 gigaohms",
	},
	{
		description: "black and red and black",
		input:       []string{"black", "red", "black"},
		expected:    "2 ohms",
	},
	{
		description: "black and red and brown",
		input:       []string{"black", "red", "brown"},
		expected:    "20 ohms",
	},
	{
		description: "black and red and red",
		input:       []string{"black", "red", "red"},
		expected:    "200 ohms",
	},
	{
		description: "black and red and orange",
		input:       []string{"black", "red", "orange"},
		expected:    "2 kiloohms",
	},
	{
		description: "black and red and yellow",
		input:       []string{"black", "red", "yellow"},
		expected:    "20 kiloohms",
	},
	{
		description: "black and red and green",
		input:       []string{"black", "red", "green"},
		expected:    "200 kiloohms",
	},
	{
		description: "black and red and blue",
		input:       []string{"black", "red", "blue"},
		expected:    "2 megaohms",
	},
	{
		description: "black and red and violet",
		input:       []string{"black", "red", "violet"},
		expected:    "20 megaohms",
	},
	{
		description: "black and red and grey",
		input:       []string{"black", "red", "grey"},
		expected:    "200 megaohms",
	},
	{
		description: "black and red and white",
		input:       []string{"black", "red", "white"},
		expected:    "2 gigaohms",
	},
	{
		description: "black and orange and black",
		input:       []string{"black", "orange", "black"},
		expected:    "3 ohms",
	},
	{
		description: "black and orange and brown",
		input:       []string{"black", "orange", "brown"},
		expected:    "30 ohms",
	},
	{
		description: "black and orange and red",
		input:       []string{"black", "orange", "red"},
		expected:    "300 ohms",
	},
	{
		description: "black and orange and orange",
		input:       []string{"black", "orange", "orange"},
		expected:    "3 kiloohms",
	},
	{
		description: "black and orange and yellow",
		input:       []string{"black", "orange", "yellow"},
		expected:    "30 kiloohms",
	},
	{
		description: "black and orange and green",
		input:       []string{"black", "orange", "green"},
		expected:    "300 kiloohms",
	},
	{
		description: "black and orange and blue",
		input:       []string{"black", "orange", "blue"},
		expected:    "3 megaohms",
	},
	{
		description: "black and orange and violet",
		input:       []string{"black", "orange", "violet"},
		expected:    "30 megaohms",
	},
	{
		description: "black and orange and grey",
		input:       []string{"black", "orange", "grey"},
		expected:    "300 megaohms",
	},
	{
		description: "black and orange and white",
		input:       []string{"black", "orange", "white"},
		expected:    "3 gigaohms",
	},
	{
		description: "black and yellow and black",
		input:       []string{"black", "yellow", "black"},
		expected:    "4 ohms",
	},
	{
		description: "black and yellow and brown",
		input:       []string{"black", "yellow", "brown"},
		expected:    "40 ohms",
	},
	{
		description: "black and yellow and red",
		input:       []string{"black", "yellow", "red"},
		expected:    "400 ohms",
	},
	{
		description: "black and yellow and orange",
		input:       []string{"black", "yellow", "orange"},
		expected:    "4 kiloohms",
	},
	{
		description: "black and yellow and yellow",
		input:       []string{"black", "yellow", "yellow"},
		expected:    "40 kiloohms",
	},
	{
		description: "black and yellow and green",
		input:       []string{"black", "yellow", "green"},
		expected:    "400 kiloohms",
	},
	{
		description: "black and yellow and blue",
		input:       []string{"black", "yellow", "blue"},
		expected:    "4 megaohms",
	},
	{
		description: "black and yellow and violet",
		input:       []string{"black", "yellow", "violet"},
		expected:    "40 megaohms",
	},
	{
		description: "black and yellow and grey",
		input:       []string{"black", "yellow", "grey"},
		expected:    "400 megaohms",
	},
	{
		description: "black and yellow and white",
		input:       []string{"black", "yellow", "white"},
		expected:    "4 gigaohms",
	},
	{
		description: "black and green and black",
		input:       []string{"black", "green", "black"},
		expected:    "5 ohms",
	},
	{
		description: "black and green and brown",
		input:       []string{"black", "green", "brown"},
		expected:    "50 ohms",
	},
	{
		description: "black and green and red",
		input:       []string{"black", "green", "red"},
		expected:    "500 ohms",
	},
	{
		description: "black and green and orange",
		input:       []string{"black", "green", "orange"},
		expected:    "5 kiloohms",
	},
	{
		description: "black and green and yellow",
		input:       []string{"black", "green", "yellow"},
		expected:    "50 kiloohms",
	},
	{
		description: "black and green and green",
		input:       []string{"black", "green", "green"},
		expected:    "500 kiloohms",
	},
	{
		description: "black and green and blue",
		input:       []string{"black", "green", "blue"},
		expected:    "5 megaohms",
	},
	{
		description: "black and green and violet",
		input:       []string{"black", "green", "violet"},
		expected:    "50 megaohms",
	},
	{
		description: "black and green and grey",
		input:       []string{"black", "green", "grey"},
		expected:    "500 megaohms",
	},
	{
		description: "black and green and white",
		input:       []string{"black", "green", "white"},
		expected:    "5 gigaohms",
	},
	{
		description: "black and blue and black",
		input:       []string{"black", "blue", "black"},
		expected:    "6 ohms",
	},
	{
		description: "black and blue and brown",
		input:       []string{"black", "blue", "brown"},
		expected:    "60 ohms",
	},
	{
		description: "black and blue and red",
		input:       []string{"black", "blue", "red"},
		expected:    "600 ohms",
	},
	{
		description: "black and blue and orange",
		input:       []string{"black", "blue", "orange"},
		expected:    "6 kiloohms",
	},
	{
		description: "black and blue and yellow",
		input:       []string{"black", "blue", "yellow"},
		expected:    "60 kiloohms",
	},
	{
		description: "black and blue and green",
		input:       []string{"black", "blue", "green"},
		expected:    "600 kiloohms",
	},
	{
		description: "black and blue and blue",
		input:       []string{"black", "blue", "blue"},
		expected:    "6 megaohms",
	},
	{
		description: "black and blue and violet",
		input:       []string{"black", "blue", "violet"},
		expected:    "60 megaohms",
	},
	{
		description: "black and blue and grey",
		input:       []string{"black", "blue", "grey"},
		expected:    "600 megaohms",
	},
	{
		description: "black and blue and white",
		input:       []string{"black", "blue", "white"},
		expected:    "6 gigaohms",
	},
	{
		description: "black and violet and black",
		input:       []string{"black", "violet", "black"},
		expected:    "7 ohms",
	},
	{
		description: "black and violet and brown",
		input:       []string{"black", "violet", "brown"},
		expected:    "70 ohms",
	},
	{
		description: "black and violet and red",
		input:       []string{"black", "violet", "red"},
		expected:    "700 ohms",
	},
	{
		description: "black and violet and orange",
		input:       []string{"black", "violet", "orange"},
		expected:    "7 kiloohms",
	},
	{
		description: "black and violet and yellow",
		input:       []string{"black", "violet", "yellow"},
		expected:    "70 kiloohms",
	},
	{
		description: "black and violet and green",
		input:       []string{"black", "violet", "green"},
		expected:    "700 kiloohms",
	},
	{
		description: "black and violet and blue",
		input:       []string{"black", "violet", "blue"},
		expected:    "7 megaohms",
	},
	{
		description: "black and violet and violet",
		input:       []string{"black", "violet", "violet"},
		expected:    "70 megaohms",
	},
	{
		description: "black and violet and grey",
		input:       []string{"black", "violet", "grey"},
		expected:    "700 megaohms",
	},
	{
		description: "black and violet and white",
		input:       []string{"black", "violet", "white"},
		expected:    "7 gigaohms",
	},
	{
		description: "black and grey and black",
		input:       []string{"black", "grey", "black"},
		expected:    "8 ohms",
	},
	{
		description: "black and grey and brown",
		input:       []string{"black", "grey", "brown"},
		expected:    "80 ohms",
	},
	{
		description: "black and grey and red",
		input:       []string{"black", "grey", "red"},
		expected:    "800 ohms",
	},
	{
		description: "black and grey and orange",
		input:       []string{"black", "grey", "orange"},
		expected:    "8 kiloohms",
	},
	{
		description: "black and grey and yellow",
		input:       []string{"black", "grey", "yellow"},
		expected:    "80 kiloohms",
	},
	{
		description: "black and grey and green",
		input:       []string{"black", "grey", "green"},
		expected:    "800 kiloohms",
	},
	{
		description: "black and grey and blue",
		input:       []string{"black", "grey", "blue"},
		expected:    "8 megaohms",
	},
	{
		description: "black and grey and violet",
		input:       []string{"black", "grey", "violet"},
		expected:    "80 megaohms",
	},
	{
		description: "black and grey and grey",
		input:       []string{"black", "grey", "grey"},
		expected:    "800 megaohms",
	},
	{
		description: "black and grey and white",
		input:       []string{"black", "grey", "white"},
		expected:    "8 gigaohms",
	},
	{
		description: "black and white and black",
		input:       []string{"black", "white", "black"},
		expected:    "9 ohms",
	},
	{
		description: "black and white and brown",
		input:       []string{"black", "white", "brown"},
		expected:    "90 ohms",
	},
	{
		description: "black and white and red",
		input:       []string{"black", "white", "red"},
		expected:    "900 ohms",
	},
	{
		description: "black and white and orange",
		input:       []string{"black", "white", "orange"},
		expected:    "9 kiloohms",
	},
	{
		description: "black and white and yellow",
		input:       []string{"black", "white", "yellow"},
		expected:    "90 kiloohms",
	},
	{
		description: "black and white and green",
		input:       []string{"black", "white", "green"},
		expected:    "900 kiloohms",
	},
	{
		description: "black and white and blue",
		input:       []string{"black", "white", "blue"},
		expected:    "9 megaohms",
	},
	{
		description: "black and white and violet",
		input:       []string{"black", "white", "violet"},
		expected:    "90 megaohms",
	},
	{
		description: "black and white and grey",
		input:       []string{"black", "white", "grey"},
		expected:    "900 megaohms",
	},
	{
		description: "black and white and white",
		input:       []string{"black", "white", "white"},
		expected:    "9 gigaohms",
	},
	{
		description: "brown and black and black",
		input:       []string{"brown", "black", "black"},
		expected:    "10 ohms",
	},
	{
		description: "brown and black and brown",
		input:       []string{"brown", "black", "brown"},
		expected:    "100 ohms",
	},
	{
		description: "brown and black and red",
		input:       []string{"brown", "black", "red"},
		expected:    "1 kiloohms",
	},
	{
		description: "brown and black and orange",
		input:       []string{"brown", "black", "orange"},
		expected:    "10 kiloohms",
	},
	{
		description: "brown and black and yellow",
		input:       []string{"brown", "black", "yellow"},
		expected:    "100 kiloohms",
	},
	{
		description: "brown and black and green",
		input:       []string{"brown", "black", "green"},
		expected:    "1 megaohms",
	},
	{
		description: "brown and black and blue",
		input:       []string{"brown", "black", "blue"},
		expected:    "10 megaohms",
	},
	{
		description: "brown and black and violet",
		input:       []string{"brown", "black", "violet"},
		expected:    "100 megaohms",
	},
	{
		description: "brown and black and grey",
		input:       []string{"brown", "black", "grey"},
		expected:    "1 gigaohms",
	},
	{
		description: "brown and black and white",
		input:       []string{"brown", "black", "white"},
		expected:    "10 gigaohms",
	},
	{
		description: "brown and brown and black",
		input:       []string{"brown", "brown", "black"},
		expected:    "11 ohms",
	},
	{
		description: "brown and brown and brown",
		input:       []string{"brown", "brown", "brown"},
		expected:    "110 ohms",
	},
	{
		description: "brown and brown and red",
		input:       []string{"brown", "brown", "red"},
		expected:    "1100 ohms",
	},
	{
		description: "brown and brown and orange",
		input:       []string{"brown", "brown", "orange"},
		expected:    "11 kiloohms",
	},
	{
		description: "brown and brown and yellow",
		input:       []string{"brown", "brown", "yellow"},
		expected:    "110 kiloohms",
	},
	{
		description: "brown and brown and green",
		input:       []string{"brown", "brown", "green"},
		expected:    "1100 kiloohms",
	},
	{
		description: "brown and brown and blue",
		input:       []string{"brown", "brown", "blue"},
		expected:    "11 megaohms",
	},
	{
		description: "brown and brown and violet",
		input:       []string{"brown", "brown", "violet"},
		expected:    "110 megaohms",
	},
	{
		description: "brown and brown and grey",
		input:       []string{"brown", "brown", "grey"},
		expected:    "1100 megaohms",
	},
	{
		description: "brown and brown and white",
		input:       []string{"brown", "brown", "white"},
		expected:    "11 gigaohms",
	},
	{
		description: "brown and red and black",
		input:       []string{"brown", "red", "black"},
		expected:    "12 ohms",
	},
	{
		description: "brown and red and brown",
		input:       []string{"brown", "red", "brown"},
		expected:    "120 ohms",
	},
	{
		description: "brown and red and red",
		input:       []string{"brown", "red", "red"},
		expected:    "1200 ohms",
	},
	{
		description: "brown and red and orange",
		input:       []string{"brown", "red", "orange"},
		expected:    "12 kiloohms",
	},
	{
		description: "brown and red and yellow",
		input:       []string{"brown", "red", "yellow"},
		expected:    "120 kiloohms",
	},
	{
		description: "brown and red and green",
		input:       []string{"brown", "red", "green"},
		expected:    "1200 kiloohms",
	},
	{
		description: "brown and red and blue",
		input:       []string{"brown", "red", "blue"},
		expected:    "12 megaohms",
	},
	{
		description: "brown and red and violet",
		input:       []string{"brown", "red", "violet"},
		expected:    "120 megaohms",
	},
	{
		description: "brown and red and grey",
		input:       []string{"brown", "red", "grey"},
		expected:    "1200 megaohms",
	},
	{
		description: "brown and red and white",
		input:       []string{"brown", "red", "white"},
		expected:    "12 gigaohms",
	},
	{
		description: "brown and orange and black",
		input:       []string{"brown", "orange", "black"},
		expected:    "13 ohms",
	},
	{
		description: "brown and orange and brown",
		input:       []string{"brown", "orange", "brown"},
		expected:    "130 ohms",
	},
	{
		description: "brown and orange and red",
		input:       []string{"brown", "orange", "red"},
		expected:    "1300 ohms",
	},
	{
		description: "brown and orange and orange",
		input:       []string{"brown", "orange", "orange"},
		expected:    "13 kiloohms",
	},
	{
		description: "brown and orange and yellow",
		input:       []string{"brown", "orange", "yellow"},
		expected:    "130 kiloohms",
	},
	{
		description: "brown and orange and green",
		input:       []string{"brown", "orange", "green"},
		expected:    "1300 kiloohms",
	},
	{
		description: "brown and orange and blue",
		input:       []string{"brown", "orange", "blue"},
		expected:    "13 megaohms",
	},
	{
		description: "brown and orange and violet",
		input:       []string{"brown", "orange", "violet"},
		expected:    "130 megaohms",
	},
	{
		description: "brown and orange and grey",
		input:       []string{"brown", "orange", "grey"},
		expected:    "1300 megaohms",
	},
	{
		description: "brown and orange and white",
		input:       []string{"brown", "orange", "white"},
		expected:    "13 gigaohms",
	},
	{
		description: "brown and yellow and black",
		input:       []string{"brown", "yellow", "black"},
		expected:    "14 ohms",
	},
	{
		description: "brown and yellow and brown",
		input:       []string{"brown", "yellow", "brown"},
		expected:    "140 ohms",
	},
	{
		description: "brown and yellow and red",
		input:       []string{"brown", "yellow", "red"},
		expected:    "1400 ohms",
	},
	{
		description: "brown and yellow and orange",
		input:       []string{"brown", "yellow", "orange"},
		expected:    "14 kiloohms",
	},
	{
		description: "brown and yellow and yellow",
		input:       []string{"brown", "yellow", "yellow"},
		expected:    "140 kiloohms",
	},
	{
		description: "brown and yellow and green",
		input:       []string{"brown", "yellow", "green"},
		expected:    "1400 kiloohms",
	},
	{
		description: "brown and yellow and blue",
		input:       []string{"brown", "yellow", "blue"},
		expected:    "14 megaohms",
	},
	{
		description: "brown and yellow and violet",
		input:       []string{"brown", "yellow", "violet"},
		expected:    "140 megaohms",
	},
	{
		description: "brown and yellow and grey",
		input:       []string{"brown", "yellow", "grey"},
		expected:    "1400 megaohms",
	},
	{
		description: "brown and yellow and white",
		input:       []string{"brown", "yellow", "white"},
		expected:    "14 gigaohms",
	},
	{
		description: "brown and green and black",
		input:       []string{"brown", "green", "black"},
		expected:    "15 ohms",
	},
	{
		description: "brown and green and brown",
		input:       []string{"brown", "green", "brown"},
		expected:    "150 ohms",
	},
	{
		description: "brown and green and red",
		input:       []string{"brown", "green", "red"},
		expected:    "1500 ohms",
	},
	{
		description: "brown and green and orange",
		input:       []string{"brown", "green", "orange"},
		expected:    "15 kiloohms",
	},
	{
		description: "brown and green and yellow",
		input:       []string{"brown", "green", "yellow"},
		expected:    "150 kiloohms",
	},
	{
		description: "brown and green and green",
		input:       []string{"brown", "green", "green"},
		expected:    "1500 kiloohms",
	},
	{
		description: "brown and green and blue",
		input:       []string{"brown", "green", "blue"},
		expected:    "15 megaohms",
	},
	{
		description: "brown and green and violet",
		input:       []string{"brown", "green", "violet"},
		expected:    "150 megaohms",
	},
	{
		description: "brown and green and grey",
		input:       []string{"brown", "green", "grey"},
		expected:    "1500 megaohms",
	},
	{
		description: "brown and green and white",
		input:       []string{"brown", "green", "white"},
		expected:    "15 gigaohms",
	},
	{
		description: "brown and blue and black",
		input:       []string{"brown", "blue", "black"},
		expected:    "16 ohms",
	},
	{
		description: "brown and blue and brown",
		input:       []string{"brown", "blue", "brown"},
		expected:    "160 ohms",
	},
	{
		description: "brown and blue and red",
		input:       []string{"brown", "blue", "red"},
		expected:    "1600 ohms",
	},
	{
		description: "brown and blue and orange",
		input:       []string{"brown", "blue", "orange"},
		expected:    "16 kiloohms",
	},
	{
		description: "brown and blue and yellow",
		input:       []string{"brown", "blue", "yellow"},
		expected:    "160 kiloohms",
	},
	{
		description: "brown and blue and green",
		input:       []string{"brown", "blue", "green"},
		expected:    "1600 kiloohms",
	},
	{
		description: "brown and blue and blue",
		input:       []string{"brown", "blue", "blue"},
		expected:    "16 megaohms",
	},
	{
		description: "brown and blue and violet",
		input:       []string{"brown", "blue", "violet"},
		expected:    "160 megaohms",
	},
	{
		description: "brown and blue and grey",
		input:       []string{"brown", "blue", "grey"},
		expected:    "1600 megaohms",
	},
	{
		description: "brown and blue and white",
		input:       []string{"brown", "blue", "white"},
		expected:    "16 gigaohms",
	},
	{
		description: "brown and violet and black",
		input:       []string{"brown", "violet", "black"},
		expected:    "17 ohms",
	},
	{
		description: "brown and violet and brown",
		input:       []string{"brown", "violet", "brown"},
		expected:    "170 ohms",
	},
	{
		description: "brown and violet and red",
		input:       []string{"brown", "violet", "red"},
		expected:    "1700 ohms",
	},
	{
		description: "brown and violet and orange",
		input:       []string{"brown", "violet", "orange"},
		expected:    "17 kiloohms",
	},
	{
		description: "brown and violet and yellow",
		input:       []string{"brown", "violet", "yellow"},
		expected:    "170 kiloohms",
	},
	{
		description: "brown and violet and green",
		input:       []string{"brown", "violet", "green"},
		expected:    "1700 kiloohms",
	},
	{
		description: "brown and violet and blue",
		input:       []string{"brown", "violet", "blue"},
		expected:    "17 megaohms",
	},
	{
		description: "brown and violet and violet",
		input:       []string{"brown", "violet", "violet"},
		expected:    "170 megaohms",
	},
	{
		description: "brown and violet and grey",
		input:       []string{"brown", "violet", "grey"},
		expected:    "1700 megaohms",
	},
	{
		description: "brown and violet and white",
		input:       []string{"brown", "violet", "white"},
		expected:    "17 gigaohms",
	},
	{
		description: "brown and grey and black",
		input:       []string{"brown", "grey", "black"},
		expected:    "18 ohms",
	},
	{
		description: "brown and grey and brown",
		input:       []string{"brown", "grey", "brown"},
		expected:    "180 ohms",
	},
	{
		description: "brown and grey and red",
		input:       []string{"brown", "grey", "red"},
		expected:    "1800 ohms",
	},
	{
		description: "brown and grey and orange",
		input:       []string{"brown", "grey", "orange"},
		expected:    "18 kiloohms",
	},
	{
		description: "brown and grey and yellow",
		input:       []string{"brown", "grey", "yellow"},
		expected:    "180 kiloohms",
	},
	{
		description: "brown and grey and green",
		input:       []string{"brown", "grey", "green"},
		expected:    "1800 kiloohms",
	},
	{
		description: "brown and grey and blue",
		input:       []string{"brown", "grey", "blue"},
		expected:    "18 megaohms",
	},
	{
		description: "brown and grey and violet",
		input:       []string{"brown", "grey", "violet"},
		expected:    "180 megaohms",
	},
	{
		description: "brown and grey and grey",
		input:       []string{"brown", "grey", "grey"},
		expected:    "1800 megaohms",
	},
	{
		description: "brown and grey and white",
		input:       []string{"brown", "grey", "white"},
		expected:    "18 gigaohms",
	},
	{
		description: "brown and white and black",
		input:       []string{"brown", "white", "black"},
		expected:    "19 ohms",
	},
	{
		description: "brown and white and brown",
		input:       []string{"brown", "white", "brown"},
		expected:    "190 ohms",
	},
	{
		description: "brown and white and red",
		input:       []string{"brown", "white", "red"},
		expected:    "1900 ohms",
	},
	{
		description: "brown and white and orange",
		input:       []string{"brown", "white", "orange"},
		expected:    "19 kiloohms",
	},
	{
		description: "brown and white and yellow",
		input:       []string{"brown", "white", "yellow"},
		expected:    "190 kiloohms",
	},
	{
		description: "brown and white and green",
		input:       []string{"brown", "white", "green"},
		expected:    "1900 kiloohms",
	},
	{
		description: "brown and white and blue",
		input:       []string{"brown", "white", "blue"},
		expected:    "19 megaohms",
	},
	{
		description: "brown and white and violet",
		input:       []string{"brown", "white", "violet"},
		expected:    "190 megaohms",
	},
	{
		description: "brown and white and grey",
		input:       []string{"brown", "white", "grey"},
		expected:    "1900 megaohms",
	},
	{
		description: "brown and white and white",
		input:       []string{"brown", "white", "white"},
		expected:    "19 gigaohms",
	},
	{
		description: "red and black and black",
		input:       []string{"red", "black", "black"},
		expected:    "20 ohms",
	},
	{
		description: "red and black and brown",
		input:       []string{"red", "black", "brown"},
		expected:    "200 ohms",
	},
	{
		description: "red and black and red",
		input:       []string{"red", "black", "red"},
		expected:    "2 kiloohms",
	},
	{
		description: "red and black and orange",
		input:       []string{"red", "black", "orange"},
		expected:    "20 kiloohms",
	},
	{
		description: "red and black and yellow",
		input:       []string{"red", "black", "yellow"},
		expected:    "200 kiloohms",
	},
	{
		description: "red and black and green",
		input:       []string{"red", "black", "green"},
		expected:    "2 megaohms",
	},
	{
		description: "red and black and blue",
		input:       []string{"red", "black", "blue"},
		expected:    "20 megaohms",
	},
	{
		description: "red and black and violet",
		input:       []string{"red", "black", "violet"},
		expected:    "200 megaohms",
	},
	{
		description: "red and black and grey",
		input:       []string{"red", "black", "grey"},
		expected:    "2 gigaohms",
	},
	{
		description: "red and black and white",
		input:       []string{"red", "black", "white"},
		expected:    "20 gigaohms",
	},
	{
		description: "red and brown and black",
		input:       []string{"red", "brown", "black"},
		expected:    "21 ohms",
	},
	{
		description: "red and brown and brown",
		input:       []string{"red", "brown", "brown"},
		expected:    "210 ohms",
	},
	{
		description: "red and brown and red",
		input:       []string{"red", "brown", "red"},
		expected:    "2100 ohms",
	},
	{
		description: "red and brown and orange",
		input:       []string{"red", "brown", "orange"},
		expected:    "21 kiloohms",
	},
	{
		description: "red and brown and yellow",
		input:       []string{"red", "brown", "yellow"},
		expected:    "210 kiloohms",
	},
	{
		description: "red and brown and green",
		input:       []string{"red", "brown", "green"},
		expected:    "2100 kiloohms",
	},
	{
		description: "red and brown and blue",
		input:       []string{"red", "brown", "blue"},
		expected:    "21 megaohms",
	},
	{
		description: "red and brown and violet",
		input:       []string{"red", "brown", "violet"},
		expected:    "210 megaohms",
	},
	{
		description: "red and brown and grey",
		input:       []string{"red", "brown", "grey"},
		expected:    "2100 megaohms",
	},
	{
		description: "red and brown and white",
		input:       []string{"red", "brown", "white"},
		expected:    "21 gigaohms",
	},
	{
		description: "red and red and black",
		input:       []string{"red", "red", "black"},
		expected:    "22 ohms",
	},
	{
		description: "red and red and brown",
		input:       []string{"red", "red", "brown"},
		expected:    "220 ohms",
	},
	{
		description: "red and red and red",
		input:       []string{"red", "red", "red"},
		expected:    "2200 ohms",
	},
	{
		description: "red and red and orange",
		input:       []string{"red", "red", "orange"},
		expected:    "22 kiloohms",
	},
	{
		description: "red and red and yellow",
		input:       []string{"red", "red", "yellow"},
		expected:    "220 kiloohms",
	},
	{
		description: "red and red and green",
		input:       []string{"red", "red", "green"},
		expected:    "2200 kiloohms",
	},
	{
		description: "red and red and blue",
		input:       []string{"red", "red", "blue"},
		expected:    "22 megaohms",
	},
	{
		description: "red and red and violet",
		input:       []string{"red", "red", "violet"},
		expected:    "220 megaohms",
	},
	{
		description: "red and red and grey",
		input:       []string{"red", "red", "grey"},
		expected:    "2200 megaohms",
	},
	{
		description: "red and red and white",
		input:       []string{"red", "red", "white"},
		expected:    "22 gigaohms",
	},
	{
		description: "red and orange and black",
		input:       []string{"red", "orange", "black"},
		expected:    "23 ohms",
	},
	{
		description: "red and orange and brown",
		input:       []string{"red", "orange", "brown"},
		expected:    "230 ohms",
	},
	{
		description: "red and orange and red",
		input:       []string{"red", "orange", "red"},
		expected:    "2300 ohms",
	},
	{
		description: "red and orange and orange",
		input:       []string{"red", "orange", "orange"},
		expected:    "23 kiloohms",
	},
	{
		description: "red and orange and yellow",
		input:       []string{"red", "orange", "yellow"},
		expected:    "230 kiloohms",
	},
	{
		description: "red and orange and green",
		input:       []string{"red", "orange", "green"},
		expected:    "2300 kiloohms",
	},
	{
		description: "red and orange and blue",
		input:       []string{"red", "orange", "blue"},
		expected:    "23 megaohms",
	},
	{
		description: "red and orange and violet",
		input:       []string{"red", "orange", "violet"},
		expected:    "230 megaohms",
	},
	{
		description: "red and orange and grey",
		input:       []string{"red", "orange", "grey"},
		expected:    "2300 megaohms",
	},
	{
		description: "red and orange and white",
		input:       []string{"red", "orange", "white"},
		expected:    "23 gigaohms",
	},
	{
		description: "red and yellow and black",
		input:       []string{"red", "yellow", "black"},
		expected:    "24 ohms",
	},
	{
		description: "red and yellow and brown",
		input:       []string{"red", "yellow", "brown"},
		expected:    "240 ohms",
	},
	{
		description: "red and yellow and red",
		input:       []string{"red", "yellow", "red"},
		expected:    "2400 ohms",
	},
	{
		description: "red and yellow and orange",
		input:       []string{"red", "yellow", "orange"},
		expected:    "24 kiloohms",
	},
	{
		description: "red and yellow and yellow",
		input:       []string{"red", "yellow", "yellow"},
		expected:    "240 kiloohms",
	},
	{
		description: "red and yellow and green",
		input:       []string{"red", "yellow", "green"},
		expected:    "2400 kiloohms",
	},
	{
		description: "red and yellow and blue",
		input:       []string{"red", "yellow", "blue"},
		expected:    "24 megaohms",
	},
	{
		description: "red and yellow and violet",
		input:       []string{"red", "yellow", "violet"},
		expected:    "240 megaohms",
	},
	{
		description: "red and yellow and grey",
		input:       []string{"red", "yellow", "grey"},
		expected:    "2400 megaohms",
	},
	{
		description: "red and yellow and white",
		input:       []string{"red", "yellow", "white"},
		expected:    "24 gigaohms",
	},
	{
		description: "red and green and black",
		input:       []string{"red", "green", "black"},
		expected:    "25 ohms",
	},
	{
		description: "red and green and brown",
		input:       []string{"red", "green", "brown"},
		expected:    "250 ohms",
	},
	{
		description: "red and green and red",
		input:       []string{"red", "green", "red"},
		expected:    "2500 ohms",
	},
	{
		description: "red and green and orange",
		input:       []string{"red", "green", "orange"},
		expected:    "25 kiloohms",
	},
	{
		description: "red and green and yellow",
		input:       []string{"red", "green", "yellow"},
		expected:    "250 kiloohms",
	},
	{
		description: "red and green and green",
		input:       []string{"red", "green", "green"},
		expected:    "2500 kiloohms",
	},
	{
		description: "red and green and blue",
		input:       []string{"red", "green", "blue"},
		expected:    "25 megaohms",
	},
	{
		description: "red and green and violet",
		input:       []string{"red", "green", "violet"},
		expected:    "250 megaohms",
	},
	{
		description: "red and green and grey",
		input:       []string{"red", "green", "grey"},
		expected:    "2500 megaohms",
	},
	{
		description: "red and green and white",
		input:       []string{"red", "green", "white"},
		expected:    "25 gigaohms",
	},
	{
		description: "red and blue and black",
		input:       []string{"red", "blue", "black"},
		expected:    "26 ohms",
	},
	{
		description: "red and blue and brown",
		input:       []string{"red", "blue", "brown"},
		expected:    "260 ohms",
	},
	{
		description: "red and blue and red",
		input:       []string{"red", "blue", "red"},
		expected:    "2600 ohms",
	},
	{
		description: "red and blue and orange",
		input:       []string{"red", "blue", "orange"},
		expected:    "26 kiloohms",
	},
	{
		description: "red and blue and yellow",
		input:       []string{"red", "blue", "yellow"},
		expected:    "260 kiloohms",
	},
	{
		description: "red and blue and green",
		input:       []string{"red", "blue", "green"},
		expected:    "2600 kiloohms",
	},
	{
		description: "red and blue and blue",
		input:       []string{"red", "blue", "blue"},
		expected:    "26 megaohms",
	},
	{
		description: "red and blue and violet",
		input:       []string{"red", "blue", "violet"},
		expected:    "260 megaohms",
	},
	{
		description: "red and blue and grey",
		input:       []string{"red", "blue", "grey"},
		expected:    "2600 megaohms",
	},
	{
		description: "red and blue and white",
		input:       []string{"red", "blue", "white"},
		expected:    "26 gigaohms",
	},
	{
		description: "red and violet and black",
		input:       []string{"red", "violet", "black"},
		expected:    "27 ohms",
	},
	{
		description: "red and violet and brown",
		input:       []string{"red", "violet", "brown"},
		expected:    "270 ohms",
	},
	{
		description: "red and violet and red",
		input:       []string{"red", "violet", "red"},
		expected:    "2700 ohms",
	},
	{
		description: "red and violet and orange",
		input:       []string{"red", "violet", "orange"},
		expected:    "27 kiloohms",
	},
	{
		description: "red and violet and yellow",
		input:       []string{"red", "violet", "yellow"},
		expected:    "270 kiloohms",
	},
	{
		description: "red and violet and green",
		input:       []string{"red", "violet", "green"},
		expected:    "2700 kiloohms",
	},
	{
		description: "red and violet and blue",
		input:       []string{"red", "violet", "blue"},
		expected:    "27 megaohms",
	},
	{
		description: "red and violet and violet",
		input:       []string{"red", "violet", "violet"},
		expected:    "270 megaohms",
	},
	{
		description: "red and violet and grey",
		input:       []string{"red", "violet", "grey"},
		expected:    "2700 megaohms",
	},
	{
		description: "red and violet and white",
		input:       []string{"red", "violet", "white"},
		expected:    "27 gigaohms",
	},
	{
		description: "red and grey and black",
		input:       []string{"red", "grey", "black"},
		expected:    "28 ohms",
	},
	{
		description: "red and grey and brown",
		input:       []string{"red", "grey", "brown"},
		expected:    "280 ohms",
	},
	{
		description: "red and grey and red",
		input:       []string{"red", "grey", "red"},
		expected:    "2800 ohms",
	},
	{
		description: "red and grey and orange",
		input:       []string{"red", "grey", "orange"},
		expected:    "28 kiloohms",
	},
	{
		description: "red and grey and yellow",
		input:       []string{"red", "grey", "yellow"},
		expected:    "280 kiloohms",
	},
	{
		description: "red and grey and green",
		input:       []string{"red", "grey", "green"},
		expected:    "2800 kiloohms",
	},
	{
		description: "red and grey and blue",
		input:       []string{"red", "grey", "blue"},
		expected:    "28 megaohms",
	},
	{
		description: "red and grey and violet",
		input:       []string{"red", "grey", "violet"},
		expected:    "280 megaohms",
	},
	{
		description: "red and grey and grey",
		input:       []string{"red", "grey", "grey"},
		expected:    "2800 megaohms",
	},
	{
		description: "red and grey and white",
		input:       []string{"red", "grey", "white"},
		expected:    "28 gigaohms",
	},
	{
		description: "red and white and black",
		input:       []string{"red", "white", "black"},
		expected:    "29 ohms",
	},
	{
		description: "red and white and brown",
		input:       []string{"red", "white", "brown"},
		expected:    "290 ohms",
	},
	{
		description: "red and white and red",
		input:       []string{"red", "white", "red"},
		expected:    "2900 ohms",
	},
	{
		description: "red and white and orange",
		input:       []string{"red", "white", "orange"},
		expected:    "29 kiloohms",
	},
	{
		description: "red and white and yellow",
		input:       []string{"red", "white", "yellow"},
		expected:    "290 kiloohms",
	},
	{
		description: "red and white and green",
		input:       []string{"red", "white", "green"},
		expected:    "2900 kiloohms",
	},
	{
		description: "red and white and blue",
		input:       []string{"red", "white", "blue"},
		expected:    "29 megaohms",
	},
	{
		description: "red and white and violet",
		input:       []string{"red", "white", "violet"},
		expected:    "290 megaohms",
	},
	{
		description: "red and white and grey",
		input:       []string{"red", "white", "grey"},
		expected:    "2900 megaohms",
	},
	{
		description: "red and white and white",
		input:       []string{"red", "white", "white"},
		expected:    "29 gigaohms",
	},
	{
		description: "orange and black and black",
		input:       []string{"orange", "black", "black"},
		expected:    "30 ohms",
	},
	{
		description: "orange and black and brown",
		input:       []string{"orange", "black", "brown"},
		expected:    "300 ohms",
	},
	{
		description: "orange and black and red",
		input:       []string{"orange", "black", "red"},
		expected:    "3 kiloohms",
	},
	{
		description: "orange and black and orange",
		input:       []string{"orange", "black", "orange"},
		expected:    "30 kiloohms",
	},
	{
		description: "orange and black and yellow",
		input:       []string{"orange", "black", "yellow"},
		expected:    "300 kiloohms",
	},
	{
		description: "orange and black and green",
		input:       []string{"orange", "black", "green"},
		expected:    "3 megaohms",
	},
	{
		description: "orange and black and blue",
		input:       []string{"orange", "black", "blue"},
		expected:    "30 megaohms",
	},
	{
		description: "orange and black and violet",
		input:       []string{"orange", "black", "violet"},
		expected:    "300 megaohms",
	},
	{
		description: "orange and black and grey",
		input:       []string{"orange", "black", "grey"},
		expected:    "3 gigaohms",
	},
	{
		description: "orange and black and white",
		input:       []string{"orange", "black", "white"},
		expected:    "30 gigaohms",
	},
	{
		description: "orange and brown and black",
		input:       []string{"orange", "brown", "black"},
		expected:    "31 ohms",
	},
	{
		description: "orange and brown and brown",
		input:       []string{"orange", "brown", "brown"},
		expected:    "310 ohms",
	},
	{
		description: "orange and brown and red",
		input:       []string{"orange", "brown", "red"},
		expected:    "3100 ohms",
	},
	{
		description: "orange and brown and orange",
		input:       []string{"orange", "brown", "orange"},
		expected:    "31 kiloohms",
	},
	{
		description: "orange and brown and yellow",
		input:       []string{"orange", "brown", "yellow"},
		expected:    "310 kiloohms",
	},
	{
		description: "orange and brown and green",
		input:       []string{"orange", "brown", "green"},
		expected:    "3100 kiloohms",
	},
	{
		description: "orange and brown and blue",
		input:       []string{"orange", "brown", "blue"},
		expected:    "31 megaohms",
	},
	{
		description: "orange and brown and violet",
		input:       []string{"orange", "brown", "violet"},
		expected:    "310 megaohms",
	},
	{
		description: "orange and brown and grey",
		input:       []string{"orange", "brown", "grey"},
		expected:    "3100 megaohms",
	},
	{
		description: "orange and brown and white",
		input:       []string{"orange", "brown", "white"},
		expected:    "31 gigaohms",
	},
	{
		description: "orange and red and black",
		input:       []string{"orange", "red", "black"},
		expected:    "32 ohms",
	},
	{
		description: "orange and red and brown",
		input:       []string{"orange", "red", "brown"},
		expected:    "320 ohms",
	},
	{
		description: "orange and red and red",
		input:       []string{"orange", "red", "red"},
		expected:    "3200 ohms",
	},
	{
		description: "orange and red and orange",
		input:       []string{"orange", "red", "orange"},
		expected:    "32 kiloohms",
	},
	{
		description: "orange and red and yellow",
		input:       []string{"orange", "red", "yellow"},
		expected:    "320 kiloohms",
	},
	{
		description: "orange and red and green",
		input:       []string{"orange", "red", "green"},
		expected:    "3200 kiloohms",
	},
	{
		description: "orange and red and blue",
		input:       []string{"orange", "red", "blue"},
		expected:    "32 megaohms",
	},
	{
		description: "orange and red and violet",
		input:       []string{"orange", "red", "violet"},
		expected:    "320 megaohms",
	},
	{
		description: "orange and red and grey",
		input:       []string{"orange", "red", "grey"},
		expected:    "3200 megaohms",
	},
	{
		description: "orange and red and white",
		input:       []string{"orange", "red", "white"},
		expected:    "32 gigaohms",
	},
	{
		description: "orange and orange and black",
		input:       []string{"orange", "orange", "black"},
		expected:    "33 ohms",
	},
	{
		description: "orange and orange and brown",
		input:       []string{"orange", "orange", "brown"},
		expected:    "330 ohms",
	},
	{
		description: "orange and orange and red",
		input:       []string{"orange", "orange", "red"},
		expected:    "3300 ohms",
	},
	{
		description: "orange and orange and orange",
		input:       []string{"orange", "orange", "orange"},
		expected:    "33 kiloohms",
	},
	{
		description: "orange and orange and yellow",
		input:       []string{"orange", "orange", "yellow"},
		expected:    "330 kiloohms",
	},
	{
		description: "orange and orange and green",
		input:       []string{"orange", "orange", "green"},
		expected:    "3300 kiloohms",
	},
	{
		description: "orange and orange and blue",
		input:       []string{"orange", "orange", "blue"},
		expected:    "33 megaohms",
	},
	{
		description: "orange and orange and violet",
		input:       []string{"orange", "orange", "violet"},
		expected:    "330 megaohms",
	},
	{
		description: "orange and orange and grey",
		input:       []string{"orange", "orange", "grey"},
		expected:    "3300 megaohms",
	},
	{
		description: "orange and orange and white",
		input:       []string{"orange", "orange", "white"},
		expected:    "33 gigaohms",
	},
	{
		description: "orange and yellow and black",
		input:       []string{"orange", "yellow", "black"},
		expected:    "34 ohms",
	},
	{
		description: "orange and yellow and brown",
		input:       []string{"orange", "yellow", "brown"},
		expected:    "340 ohms",
	},
	{
		description: "orange and yellow and red",
		input:       []string{"orange", "yellow", "red"},
		expected:    "3400 ohms",
	},
	{
		description: "orange and yellow and orange",
		input:       []string{"orange", "yellow", "orange"},
		expected:    "34 kiloohms",
	},
	{
		description: "orange and yellow and yellow",
		input:       []string{"orange", "yellow", "yellow"},
		expected:    "340 kiloohms",
	},
	{
		description: "orange and yellow and green",
		input:       []string{"orange", "yellow", "green"},
		expected:    "3400 kiloohms",
	},
	{
		description: "orange and yellow and blue",
		input:       []string{"orange", "yellow", "blue"},
		expected:    "34 megaohms",
	},
	{
		description: "orange and yellow and violet",
		input:       []string{"orange", "yellow", "violet"},
		expected:    "340 megaohms",
	},
	{
		description: "orange and yellow and grey",
		input:       []string{"orange", "yellow", "grey"},
		expected:    "3400 megaohms",
	},
	{
		description: "orange and yellow and white",
		input:       []string{"orange", "yellow", "white"},
		expected:    "34 gigaohms",
	},
	{
		description: "orange and green and black",
		input:       []string{"orange", "green", "black"},
		expected:    "35 ohms",
	},
	{
		description: "orange and green and brown",
		input:       []string{"orange", "green", "brown"},
		expected:    "350 ohms",
	},
	{
		description: "orange and green and red",
		input:       []string{"orange", "green", "red"},
		expected:    "3500 ohms",
	},
	{
		description: "orange and green and orange",
		input:       []string{"orange", "green", "orange"},
		expected:    "35 kiloohms",
	},
	{
		description: "orange and green and yellow",
		input:       []string{"orange", "green", "yellow"},
		expected:    "350 kiloohms",
	},
	{
		description: "orange and green and green",
		input:       []string{"orange", "green", "green"},
		expected:    "3500 kiloohms",
	},
	{
		description: "orange and green and blue",
		input:       []string{"orange", "green", "blue"},
		expected:    "35 megaohms",
	},
	{
		description: "orange and green and violet",
		input:       []string{"orange", "green", "violet"},
		expected:    "350 megaohms",
	},
	{
		description: "orange and green and grey",
		input:       []string{"orange", "green", "grey"},
		expected:    "3500 megaohms",
	},
	{
		description: "orange and green and white",
		input:       []string{"orange", "green", "white"},
		expected:    "35 gigaohms",
	},
	{
		description: "orange and blue and black",
		input:       []string{"orange", "blue", "black"},
		expected:    "36 ohms",
	},
	{
		description: "orange and blue and brown",
		input:       []string{"orange", "blue", "brown"},
		expected:    "360 ohms",
	},
	{
		description: "orange and blue and red",
		input:       []string{"orange", "blue", "red"},
		expected:    "3600 ohms",
	},
	{
		description: "orange and blue and orange",
		input:       []string{"orange", "blue", "orange"},
		expected:    "36 kiloohms",
	},
	{
		description: "orange and blue and yellow",
		input:       []string{"orange", "blue", "yellow"},
		expected:    "360 kiloohms",
	},
	{
		description: "orange and blue and green",
		input:       []string{"orange", "blue", "green"},
		expected:    "3600 kiloohms",
	},
	{
		description: "orange and blue and blue",
		input:       []string{"orange", "blue", "blue"},
		expected:    "36 megaohms",
	},
	{
		description: "orange and blue and violet",
		input:       []string{"orange", "blue", "violet"},
		expected:    "360 megaohms",
	},
	{
		description: "orange and blue and grey",
		input:       []string{"orange", "blue", "grey"},
		expected:    "3600 megaohms",
	},
	{
		description: "orange and blue and white",
		input:       []string{"orange", "blue", "white"},
		expected:    "36 gigaohms",
	},
	{
		description: "orange and violet and black",
		input:       []string{"orange", "violet", "black"},
		expected:    "37 ohms",
	},
	{
		description: "orange and violet and brown",
		input:       []string{"orange", "violet", "brown"},
		expected:    "370 ohms",
	},
	{
		description: "orange and violet and red",
		input:       []string{"orange", "violet", "red"},
		expected:    "3700 ohms",
	},
	{
		description: "orange and violet and orange",
		input:       []string{"orange", "violet", "orange"},
		expected:    "37 kiloohms",
	},
	{
		description: "orange and violet and yellow",
		input:       []string{"orange", "violet", "yellow"},
		expected:    "370 kiloohms",
	},
	{
		description: "orange and violet and green",
		input:       []string{"orange", "violet", "green"},
		expected:    "3700 kiloohms",
	},
	{
		description: "orange and violet and blue",
		input:       []string{"orange", "violet", "blue"},
		expected:    "37 megaohms",
	},
	{
		description: "orange and violet and violet",
		input:       []string{"orange", "violet", "violet"},
		expected:    "370 megaohms",
	},
	{
		description: "orange and violet and grey",
		input:       []string{"orange", "violet", "grey"},
		expected:    "3700 megaohms",
	},
	{
		description: "orange and violet and white",
		input:       []string{"orange", "violet", "white"},
		expected:    "37 gigaohms",
	},
	{
		description: "orange and grey and black",
		input:       []string{"orange", "grey", "black"},
		expected:    "38 ohms",
	},
	{
		description: "orange and grey and brown",
		input:       []string{"orange", "grey", "brown"},
		expected:    "380 ohms",
	},
	{
		description: "orange and grey and red",
		input:       []string{"orange", "grey", "red"},
		expected:    "3800 ohms",
	},
	{
		description: "orange and grey and orange",
		input:       []string{"orange", "grey", "orange"},
		expected:    "38 kiloohms",
	},
	{
		description: "orange and grey and yellow",
		input:       []string{"orange", "grey", "yellow"},
		expected:    "380 kiloohms",
	},
	{
		description: "orange and grey and green",
		input:       []string{"orange", "grey", "green"},
		expected:    "3800 kiloohms",
	},
	{
		description: "orange and grey and blue",
		input:       []string{"orange", "grey", "blue"},
		expected:    "38 megaohms",
	},
	{
		description: "orange and grey and violet",
		input:       []string{"orange", "grey", "violet"},
		expected:    "380 megaohms",
	},
	{
		description: "orange and grey and grey",
		input:       []string{"orange", "grey", "grey"},
		expected:    "3800 megaohms",
	},
	{
		description: "orange and grey and white",
		input:       []string{"orange", "grey", "white"},
		expected:    "38 gigaohms",
	},
	{
		description: "orange and white and black",
		input:       []string{"orange", "white", "black"},
		expected:    "39 ohms",
	},
	{
		description: "orange and white and brown",
		input:       []string{"orange", "white", "brown"},
		expected:    "390 ohms",
	},
	{
		description: "orange and white and red",
		input:       []string{"orange", "white", "red"},
		expected:    "3900 ohms",
	},
	{
		description: "orange and white and orange",
		input:       []string{"orange", "white", "orange"},
		expected:    "39 kiloohms",
	},
	{
		description: "orange and white and yellow",
		input:       []string{"orange", "white", "yellow"},
		expected:    "390 kiloohms",
	},
	{
		description: "orange and white and green",
		input:       []string{"orange", "white", "green"},
		expected:    "3900 kiloohms",
	},
	{
		description: "orange and white and blue",
		input:       []string{"orange", "white", "blue"},
		expected:    "39 megaohms",
	},
	{
		description: "orange and white and violet",
		input:       []string{"orange", "white", "violet"},
		expected:    "390 megaohms",
	},
	{
		description: "orange and white and grey",
		input:       []string{"orange", "white", "grey"},
		expected:    "3900 megaohms",
	},
	{
		description: "orange and white and white",
		input:       []string{"orange", "white", "white"},
		expected:    "39 gigaohms",
	},
	{
		description: "yellow and black and black",
		input:       []string{"yellow", "black", "black"},
		expected:    "40 ohms",
	},
	{
		description: "yellow and black and brown",
		input:       []string{"yellow", "black", "brown"},
		expected:    "400 ohms",
	},
	{
		description: "yellow and black and red",
		input:       []string{"yellow", "black", "red"},
		expected:    "4 kiloohms",
	},
	{
		description: "yellow and black and orange",
		input:       []string{"yellow", "black", "orange"},
		expected:    "40 kiloohms",
	},
	{
		description: "yellow and black and yellow",
		input:       []string{"yellow", "black", "yellow"},
		expected:    "400 kiloohms",
	},
	{
		description: "yellow and black and green",
		input:       []string{"yellow", "black", "green"},
		expected:    "4 megaohms",
	},
	{
		description: "yellow and black and blue",
		input:       []string{"yellow", "black", "blue"},
		expected:    "40 megaohms",
	},
	{
		description: "yellow and black and violet",
		input:       []string{"yellow", "black", "violet"},
		expected:    "400 megaohms",
	},
	{
		description: "yellow and black and grey",
		input:       []string{"yellow", "black", "grey"},
		expected:    "4 gigaohms",
	},
	{
		description: "yellow and black and white",
		input:       []string{"yellow", "black", "white"},
		expected:    "40 gigaohms",
	},
	{
		description: "yellow and brown and black",
		input:       []string{"yellow", "brown", "black"},
		expected:    "41 ohms",
	},
	{
		description: "yellow and brown and brown",
		input:       []string{"yellow", "brown", "brown"},
		expected:    "410 ohms",
	},
	{
		description: "yellow and brown and red",
		input:       []string{"yellow", "brown", "red"},
		expected:    "4100 ohms",
	},
	{
		description: "yellow and brown and orange",
		input:       []string{"yellow", "brown", "orange"},
		expected:    "41 kiloohms",
	},
	{
		description: "yellow and brown and yellow",
		input:       []string{"yellow", "brown", "yellow"},
		expected:    "410 kiloohms",
	},
	{
		description: "yellow and brown and green",
		input:       []string{"yellow", "brown", "green"},
		expected:    "4100 kiloohms",
	},
	{
		description: "yellow and brown and blue",
		input:       []string{"yellow", "brown", "blue"},
		expected:    "41 megaohms",
	},
	{
		description: "yellow and brown and violet",
		input:       []string{"yellow", "brown", "violet"},
		expected:    "410 megaohms",
	},
	{
		description: "yellow and brown and grey",
		input:       []string{"yellow", "brown", "grey"},
		expected:    "4100 megaohms",
	},
	{
		description: "yellow and brown and white",
		input:       []string{"yellow", "brown", "white"},
		expected:    "41 gigaohms",
	},
	{
		description: "yellow and red and black",
		input:       []string{"yellow", "red", "black"},
		expected:    "42 ohms",
	},
	{
		description: "yellow and red and brown",
		input:       []string{"yellow", "red", "brown"},
		expected:    "420 ohms",
	},
	{
		description: "yellow and red and red",
		input:       []string{"yellow", "red", "red"},
		expected:    "4200 ohms",
	},
	{
		description: "yellow and red and orange",
		input:       []string{"yellow", "red", "orange"},
		expected:    "42 kiloohms",
	},
	{
		description: "yellow and red and yellow",
		input:       []string{"yellow", "red", "yellow"},
		expected:    "420 kiloohms",
	},
	{
		description: "yellow and red and green",
		input:       []string{"yellow", "red", "green"},
		expected:    "4200 kiloohms",
	},
	{
		description: "yellow and red and blue",
		input:       []string{"yellow", "red", "blue"},
		expected:    "42 megaohms",
	},
	{
		description: "yellow and red and violet",
		input:       []string{"yellow", "red", "violet"},
		expected:    "420 megaohms",
	},
	{
		description: "yellow and red and grey",
		input:       []string{"yellow", "red", "grey"},
		expected:    "4200 megaohms",
	},
	{
		description: "yellow and red and white",
		input:       []string{"yellow", "red", "white"},
		expected:    "42 gigaohms",
	},
	{
		description: "yellow and orange and black",
		input:       []string{"yellow", "orange", "black"},
		expected:    "43 ohms",
	},
	{
		description: "yellow and orange and brown",
		input:       []string{"yellow", "orange", "brown"},
		expected:    "430 ohms",
	},
	{
		description: "yellow and orange and red",
		input:       []string{"yellow", "orange", "red"},
		expected:    "4300 ohms",
	},
	{
		description: "yellow and orange and orange",
		input:       []string{"yellow", "orange", "orange"},
		expected:    "43 kiloohms",
	},
	{
		description: "yellow and orange and yellow",
		input:       []string{"yellow", "orange", "yellow"},
		expected:    "430 kiloohms",
	},
	{
		description: "yellow and orange and green",
		input:       []string{"yellow", "orange", "green"},
		expected:    "4300 kiloohms",
	},
	{
		description: "yellow and orange and blue",
		input:       []string{"yellow", "orange", "blue"},
		expected:    "43 megaohms",
	},
	{
		description: "yellow and orange and violet",
		input:       []string{"yellow", "orange", "violet"},
		expected:    "430 megaohms",
	},
	{
		description: "yellow and orange and grey",
		input:       []string{"yellow", "orange", "grey"},
		expected:    "4300 megaohms",
	},
	{
		description: "yellow and orange and white",
		input:       []string{"yellow", "orange", "white"},
		expected:    "43 gigaohms",
	},
	{
		description: "yellow and yellow and black",
		input:       []string{"yellow", "yellow", "black"},
		expected:    "44 ohms",
	},
	{
		description: "yellow and yellow and brown",
		input:       []string{"yellow", "yellow", "brown"},
		expected:    "440 ohms",
	},
	{
		description: "yellow and yellow and red",
		input:       []string{"yellow", "yellow", "red"},
		expected:    "4400 ohms",
	},
	{
		description: "yellow and yellow and orange",
		input:       []string{"yellow", "yellow", "orange"},
		expected:    "44 kiloohms",
	},
	{
		description: "yellow and yellow and yellow",
		input:       []string{"yellow", "yellow", "yellow"},
		expected:    "440 kiloohms",
	},
	{
		description: "yellow and yellow and green",
		input:       []string{"yellow", "yellow", "green"},
		expected:    "4400 kiloohms",
	},
	{
		description: "yellow and yellow and blue",
		input:       []string{"yellow", "yellow", "blue"},
		expected:    "44 megaohms",
	},
	{
		description: "yellow and yellow and violet",
		input:       []string{"yellow", "yellow", "violet"},
		expected:    "440 megaohms",
	},
	{
		description: "yellow and yellow and grey",
		input:       []string{"yellow", "yellow", "grey"},
		expected:    "4400 megaohms",
	},
	{
		description: "yellow and yellow and white",
		input:       []string{"yellow", "yellow", "white"},
		expected:    "44 gigaohms",
	},
	{
		description: "yellow and green and black",
		input:       []string{"yellow", "green", "black"},
		expected:    "45 ohms",
	},
	{
		description: "yellow and green and brown",
		input:       []string{"yellow", "green", "brown"},
		expected:    "450 ohms",
	},
	{
		description: "yellow and green and red",
		input:       []string{"yellow", "green", "red"},
		expected:    "4500 ohms",
	},
	{
		description: "yellow and green and orange",
		input:       []string{"yellow", "green", "orange"},
		expected:    "45 kiloohms",
	},
	{
		description: "yellow and green and yellow",
		input:       []string{"yellow", "green", "yellow"},
		expected:    "450 kiloohms",
	},
	{
		description: "yellow and green and green",
		input:       []string{"yellow", "green", "green"},
		expected:    "4500 kiloohms",
	},
	{
		description: "yellow and green and blue",
		input:       []string{"yellow", "green", "blue"},
		expected:    "45 megaohms",
	},
	{
		description: "yellow and green and violet",
		input:       []string{"yellow", "green", "violet"},
		expected:    "450 megaohms",
	},
	{
		description: "yellow and green and grey",
		input:       []string{"yellow", "green", "grey"},
		expected:    "4500 megaohms",
	},
	{
		description: "yellow and green and white",
		input:       []string{"yellow", "green", "white"},
		expected:    "45 gigaohms",
	},
	{
		description: "yellow and blue and black",
		input:       []string{"yellow", "blue", "black"},
		expected:    "46 ohms",
	},
	{
		description: "yellow and blue and brown",
		input:       []string{"yellow", "blue", "brown"},
		expected:    "460 ohms",
	},
	{
		description: "yellow and blue and red",
		input:       []string{"yellow", "blue", "red"},
		expected:    "4600 ohms",
	},
	{
		description: "yellow and blue and orange",
		input:       []string{"yellow", "blue", "orange"},
		expected:    "46 kiloohms",
	},
	{
		description: "yellow and blue and yellow",
		input:       []string{"yellow", "blue", "yellow"},
		expected:    "460 kiloohms",
	},
	{
		description: "yellow and blue and green",
		input:       []string{"yellow", "blue", "green"},
		expected:    "4600 kiloohms",
	},
	{
		description: "yellow and blue and blue",
		input:       []string{"yellow", "blue", "blue"},
		expected:    "46 megaohms",
	},
	{
		description: "yellow and blue and violet",
		input:       []string{"yellow", "blue", "violet"},
		expected:    "460 megaohms",
	},
	{
		description: "yellow and blue and grey",
		input:       []string{"yellow", "blue", "grey"},
		expected:    "4600 megaohms",
	},
	{
		description: "yellow and blue and white",
		input:       []string{"yellow", "blue", "white"},
		expected:    "46 gigaohms",
	},
	{
		description: "yellow and violet and black",
		input:       []string{"yellow", "violet", "black"},
		expected:    "47 ohms",
	},
	{
		description: "yellow and violet and brown",
		input:       []string{"yellow", "violet", "brown"},
		expected:    "470 ohms",
	},
	{
		description: "yellow and violet and red",
		input:       []string{"yellow", "violet", "red"},
		expected:    "4700 ohms",
	},
	{
		description: "yellow and violet and orange",
		input:       []string{"yellow", "violet", "orange"},
		expected:    "47 kiloohms",
	},
	{
		description: "yellow and violet and yellow",
		input:       []string{"yellow", "violet", "yellow"},
		expected:    "470 kiloohms",
	},
	{
		description: "yellow and violet and green",
		input:       []string{"yellow", "violet", "green"},
		expected:    "4700 kiloohms",
	},
	{
		description: "yellow and violet and blue",
		input:       []string{"yellow", "violet", "blue"},
		expected:    "47 megaohms",
	},
	{
		description: "yellow and violet and violet",
		input:       []string{"yellow", "violet", "violet"},
		expected:    "470 megaohms",
	},
	{
		description: "yellow and violet and grey",
		input:       []string{"yellow", "violet", "grey"},
		expected:    "4700 megaohms",
	},
	{
		description: "yellow and violet and white",
		input:       []string{"yellow", "violet", "white"},
		expected:    "47 gigaohms",
	},
	{
		description: "yellow and grey and black",
		input:       []string{"yellow", "grey", "black"},
		expected:    "48 ohms",
	},
	{
		description: "yellow and grey and brown",
		input:       []string{"yellow", "grey", "brown"},
		expected:    "480 ohms",
	},
	{
		description: "yellow and grey and red",
		input:       []string{"yellow", "grey", "red"},
		expected:    "4800 ohms",
	},
	{
		description: "yellow and grey and orange",
		input:       []string{"yellow", "grey", "orange"},
		expected:    "48 kiloohms",
	},
	{
		description: "yellow and grey and yellow",
		input:       []string{"yellow", "grey", "yellow"},
		expected:    "480 kiloohms",
	},
	{
		description: "yellow and grey and green",
		input:       []string{"yellow", "grey", "green"},
		expected:    "4800 kiloohms",
	},
	{
		description: "yellow and grey and blue",
		input:       []string{"yellow", "grey", "blue"},
		expected:    "48 megaohms",
	},
	{
		description: "yellow and grey and violet",
		input:       []string{"yellow", "grey", "violet"},
		expected:    "480 megaohms",
	},
	{
		description: "yellow and grey and grey",
		input:       []string{"yellow", "grey", "grey"},
		expected:    "4800 megaohms",
	},
	{
		description: "yellow and grey and white",
		input:       []string{"yellow", "grey", "white"},
		expected:    "48 gigaohms",
	},
	{
		description: "yellow and white and black",
		input:       []string{"yellow", "white", "black"},
		expected:    "49 ohms",
	},
	{
		description: "yellow and white and brown",
		input:       []string{"yellow", "white", "brown"},
		expected:    "490 ohms",
	},
	{
		description: "yellow and white and red",
		input:       []string{"yellow", "white", "red"},
		expected:    "4900 ohms",
	},
	{
		description: "yellow and white and orange",
		input:       []string{"yellow", "white", "orange"},
		expected:    "49 kiloohms",
	},
	{
		description: "yellow and white and yellow",
		input:       []string{"yellow", "white", "yellow"},
		expected:    "490 kiloohms",
	},
	{
		description: "yellow and white and green",
		input:       []string{"yellow", "white", "green"},
		expected:    "4900 kiloohms",
	},
	{
		description: "yellow and white and blue",
		input:       []string{"yellow", "white", "blue"},
		expected:    "49 megaohms",
	},
	{
		description: "yellow and white and violet",
		input:       []string{"yellow", "white", "violet"},
		expected:    "490 megaohms",
	},
	{
		description: "yellow and white and grey",
		input:       []string{"yellow", "white", "grey"},
		expected:    "4900 megaohms",
	},
	{
		description: "yellow and white and white",
		input:       []string{"yellow", "white", "white"},
		expected:    "49 gigaohms",
	},
	{
		description: "green and black and black",
		input:       []string{"green", "black", "black"},
		expected:    "50 ohms",
	},
	{
		description: "green and black and brown",
		input:       []string{"green", "black", "brown"},
		expected:    "500 ohms",
	},
	{
		description: "green and black and red",
		input:       []string{"green", "black", "red"},
		expected:    "5 kiloohms",
	},
	{
		description: "green and black and orange",
		input:       []string{"green", "black", "orange"},
		expected:    "50 kiloohms",
	},
	{
		description: "green and black and yellow",
		input:       []string{"green", "black", "yellow"},
		expected:    "500 kiloohms",
	},
	{
		description: "green and black and green",
		input:       []string{"green", "black", "green"},
		expected:    "5 megaohms",
	},
	{
		description: "green and black and blue",
		input:       []string{"green", "black", "blue"},
		expected:    "50 megaohms",
	},
	{
		description: "green and black and violet",
		input:       []string{"green", "black", "violet"},
		expected:    "500 megaohms",
	},
	{
		description: "green and black and grey",
		input:       []string{"green", "black", "grey"},
		expected:    "5 gigaohms",
	},
	{
		description: "green and black and white",
		input:       []string{"green", "black", "white"},
		expected:    "50 gigaohms",
	},
	{
		description: "green and brown and black",
		input:       []string{"green", "brown", "black"},
		expected:    "51 ohms",
	},
	{
		description: "green and brown and brown",
		input:       []string{"green", "brown", "brown"},
		expected:    "510 ohms",
	},
	{
		description: "green and brown and red",
		input:       []string{"green", "brown", "red"},
		expected:    "5100 ohms",
	},
	{
		description: "green and brown and orange",
		input:       []string{"green", "brown", "orange"},
		expected:    "51 kiloohms",
	},
	{
		description: "green and brown and yellow",
		input:       []string{"green", "brown", "yellow"},
		expected:    "510 kiloohms",
	},
	{
		description: "green and brown and green",
		input:       []string{"green", "brown", "green"},
		expected:    "5100 kiloohms",
	},
	{
		description: "green and brown and blue",
		input:       []string{"green", "brown", "blue"},
		expected:    "51 megaohms",
	},
	{
		description: "green and brown and violet",
		input:       []string{"green", "brown", "violet"},
		expected:    "510 megaohms",
	},
	{
		description: "green and brown and grey",
		input:       []string{"green", "brown", "grey"},
		expected:    "5100 megaohms",
	},
	{
		description: "green and brown and white",
		input:       []string{"green", "brown", "white"},
		expected:    "51 gigaohms",
	},
	{
		description: "green and red and black",
		input:       []string{"green", "red", "black"},
		expected:    "52 ohms",
	},
	{
		description: "green and red and brown",
		input:       []string{"green", "red", "brown"},
		expected:    "520 ohms",
	},
	{
		description: "green and red and red",
		input:       []string{"green", "red", "red"},
		expected:    "5200 ohms",
	},
	{
		description: "green and red and orange",
		input:       []string{"green", "red", "orange"},
		expected:    "52 kiloohms",
	},
	{
		description: "green and red and yellow",
		input:       []string{"green", "red", "yellow"},
		expected:    "520 kiloohms",
	},
	{
		description: "green and red and green",
		input:       []string{"green", "red", "green"},
		expected:    "5200 kiloohms",
	},
	{
		description: "green and red and blue",
		input:       []string{"green", "red", "blue"},
		expected:    "52 megaohms",
	},
	{
		description: "green and red and violet",
		input:       []string{"green", "red", "violet"},
		expected:    "520 megaohms",
	},
	{
		description: "green and red and grey",
		input:       []string{"green", "red", "grey"},
		expected:    "5200 megaohms",
	},
	{
		description: "green and red and white",
		input:       []string{"green", "red", "white"},
		expected:    "52 gigaohms",
	},
	{
		description: "green and orange and black",
		input:       []string{"green", "orange", "black"},
		expected:    "53 ohms",
	},
	{
		description: "green and orange and brown",
		input:       []string{"green", "orange", "brown"},
		expected:    "530 ohms",
	},
	{
		description: "green and orange and red",
		input:       []string{"green", "orange", "red"},
		expected:    "5300 ohms",
	},
	{
		description: "green and orange and orange",
		input:       []string{"green", "orange", "orange"},
		expected:    "53 kiloohms",
	},
	{
		description: "green and orange and yellow",
		input:       []string{"green", "orange", "yellow"},
		expected:    "530 kiloohms",
	},
	{
		description: "green and orange and green",
		input:       []string{"green", "orange", "green"},
		expected:    "5300 kiloohms",
	},
	{
		description: "green and orange and blue",
		input:       []string{"green", "orange", "blue"},
		expected:    "53 megaohms",
	},
	{
		description: "green and orange and violet",
		input:       []string{"green", "orange", "violet"},
		expected:    "530 megaohms",
	},
	{
		description: "green and orange and grey",
		input:       []string{"green", "orange", "grey"},
		expected:    "5300 megaohms",
	},
	{
		description: "green and orange and white",
		input:       []string{"green", "orange", "white"},
		expected:    "53 gigaohms",
	},
	{
		description: "green and yellow and black",
		input:       []string{"green", "yellow", "black"},
		expected:    "54 ohms",
	},
	{
		description: "green and yellow and brown",
		input:       []string{"green", "yellow", "brown"},
		expected:    "540 ohms",
	},
	{
		description: "green and yellow and red",
		input:       []string{"green", "yellow", "red"},
		expected:    "5400 ohms",
	},
	{
		description: "green and yellow and orange",
		input:       []string{"green", "yellow", "orange"},
		expected:    "54 kiloohms",
	},
	{
		description: "green and yellow and yellow",
		input:       []string{"green", "yellow", "yellow"},
		expected:    "540 kiloohms",
	},
	{
		description: "green and yellow and green",
		input:       []string{"green", "yellow", "green"},
		expected:    "5400 kiloohms",
	},
	{
		description: "green and yellow and blue",
		input:       []string{"green", "yellow", "blue"},
		expected:    "54 megaohms",
	},
	{
		description: "green and yellow and violet",
		input:       []string{"green", "yellow", "violet"},
		expected:    "540 megaohms",
	},
	{
		description: "green and yellow and grey",
		input:       []string{"green", "yellow", "grey"},
		expected:    "5400 megaohms",
	},
	{
		description: "green and yellow and white",
		input:       []string{"green", "yellow", "white"},
		expected:    "54 gigaohms",
	},
	{
		description: "green and green and black",
		input:       []string{"green", "green", "black"},
		expected:    "55 ohms",
	},
	{
		description: "green and green and brown",
		input:       []string{"green", "green", "brown"},
		expected:    "550 ohms",
	},
	{
		description: "green and green and red",
		input:       []string{"green", "green", "red"},
		expected:    "5500 ohms",
	},
	{
		description: "green and green and orange",
		input:       []string{"green", "green", "orange"},
		expected:    "55 kiloohms",
	},
	{
		description: "green and green and yellow",
		input:       []string{"green", "green", "yellow"},
		expected:    "550 kiloohms",
	},
	{
		description: "green and green and green",
		input:       []string{"green", "green", "green"},
		expected:    "5500 kiloohms",
	},
	{
		description: "green and green and blue",
		input:       []string{"green", "green", "blue"},
		expected:    "55 megaohms",
	},
	{
		description: "green and green and violet",
		input:       []string{"green", "green", "violet"},
		expected:    "550 megaohms",
	},
	{
		description: "green and green and grey",
		input:       []string{"green", "green", "grey"},
		expected:    "5500 megaohms",
	},
	{
		description: "green and green and white",
		input:       []string{"green", "green", "white"},
		expected:    "55 gigaohms",
	},
	{
		description: "green and blue and black",
		input:       []string{"green", "blue", "black"},
		expected:    "56 ohms",
	},
	{
		description: "green and blue and brown",
		input:       []string{"green", "blue", "brown"},
		expected:    "560 ohms",
	},
	{
		description: "green and blue and red",
		input:       []string{"green", "blue", "red"},
		expected:    "5600 ohms",
	},
	{
		description: "green and blue and orange",
		input:       []string{"green", "blue", "orange"},
		expected:    "56 kiloohms",
	},
	{
		description: "green and blue and yellow",
		input:       []string{"green", "blue", "yellow"},
		expected:    "560 kiloohms",
	},
	{
		description: "green and blue and green",
		input:       []string{"green", "blue", "green"},
		expected:    "5600 kiloohms",
	},
	{
		description: "green and blue and blue",
		input:       []string{"green", "blue", "blue"},
		expected:    "56 megaohms",
	},
	{
		description: "green and blue and violet",
		input:       []string{"green", "blue", "violet"},
		expected:    "560 megaohms",
	},
	{
		description: "green and blue and grey",
		input:       []string{"green", "blue", "grey"},
		expected:    "5600 megaohms",
	},
	{
		description: "green and blue and white",
		input:       []string{"green", "blue", "white"},
		expected:    "56 gigaohms",
	},
	{
		description: "green and violet and black",
		input:       []string{"green", "violet", "black"},
		expected:    "57 ohms",
	},
	{
		description: "green and violet and brown",
		input:       []string{"green", "violet", "brown"},
		expected:    "570 ohms",
	},
	{
		description: "green and violet and red",
		input:       []string{"green", "violet", "red"},
		expected:    "5700 ohms",
	},
	{
		description: "green and violet and orange",
		input:       []string{"green", "violet", "orange"},
		expected:    "57 kiloohms",
	},
	{
		description: "green and violet and yellow",
		input:       []string{"green", "violet", "yellow"},
		expected:    "570 kiloohms",
	},
	{
		description: "green and violet and green",
		input:       []string{"green", "violet", "green"},
		expected:    "5700 kiloohms",
	},
	{
		description: "green and violet and blue",
		input:       []string{"green", "violet", "blue"},
		expected:    "57 megaohms",
	},
	{
		description: "green and violet and violet",
		input:       []string{"green", "violet", "violet"},
		expected:    "570 megaohms",
	},
	{
		description: "green and violet and grey",
		input:       []string{"green", "violet", "grey"},
		expected:    "5700 megaohms",
	},
	{
		description: "green and violet and white",
		input:       []string{"green", "violet", "white"},
		expected:    "57 gigaohms",
	},
	{
		description: "green and grey and black",
		input:       []string{"green", "grey", "black"},
		expected:    "58 ohms",
	},
	{
		description: "green and grey and brown",
		input:       []string{"green", "grey", "brown"},
		expected:    "580 ohms",
	},
	{
		description: "green and grey and red",
		input:       []string{"green", "grey", "red"},
		expected:    "5800 ohms",
	},
	{
		description: "green and grey and orange",
		input:       []string{"green", "grey", "orange"},
		expected:    "58 kiloohms",
	},
	{
		description: "green and grey and yellow",
		input:       []string{"green", "grey", "yellow"},
		expected:    "580 kiloohms",
	},
	{
		description: "green and grey and green",
		input:       []string{"green", "grey", "green"},
		expected:    "5800 kiloohms",
	},
	{
		description: "green and grey and blue",
		input:       []string{"green", "grey", "blue"},
		expected:    "58 megaohms",
	},
	{
		description: "green and grey and violet",
		input:       []string{"green", "grey", "violet"},
		expected:    "580 megaohms",
	},
	{
		description: "green and grey and grey",
		input:       []string{"green", "grey", "grey"},
		expected:    "5800 megaohms",
	},
	{
		description: "green and grey and white",
		input:       []string{"green", "grey", "white"},
		expected:    "58 gigaohms",
	},
	{
		description: "green and white and black",
		input:       []string{"green", "white", "black"},
		expected:    "59 ohms",
	},
	{
		description: "green and white and brown",
		input:       []string{"green", "white", "brown"},
		expected:    "590 ohms",
	},
	{
		description: "green and white and red",
		input:       []string{"green", "white", "red"},
		expected:    "5900 ohms",
	},
	{
		description: "green and white and orange",
		input:       []string{"green", "white", "orange"},
		expected:    "59 kiloohms",
	},
	{
		description: "green and white and yellow",
		input:       []string{"green", "white", "yellow"},
		expected:    "590 kiloohms",
	},
	{
		description: "green and white and green",
		input:       []string{"green", "white", "green"},
		expected:    "5900 kiloohms",
	},
	{
		description: "green and white and blue",
		input:       []string{"green", "white", "blue"},
		expected:    "59 megaohms",
	},
	{
		description: "green and white and violet",
		input:       []string{"green", "white", "violet"},
		expected:    "590 megaohms",
	},
	{
		description: "green and white and grey",
		input:       []string{"green", "white", "grey"},
		expected:    "5900 megaohms",
	},
	{
		description: "green and white and white",
		input:       []string{"green", "white", "white"},
		expected:    "59 gigaohms",
	},
	{
		description: "blue and black and black",
		input:       []string{"blue", "black", "black"},
		expected:    "60 ohms",
	},
	{
		description: "blue and black and brown",
		input:       []string{"blue", "black", "brown"},
		expected:    "600 ohms",
	},
	{
		description: "blue and black and red",
		input:       []string{"blue", "black", "red"},
		expected:    "6 kiloohms",
	},
	{
		description: "blue and black and orange",
		input:       []string{"blue", "black", "orange"},
		expected:    "60 kiloohms",
	},
	{
		description: "blue and black and yellow",
		input:       []string{"blue", "black", "yellow"},
		expected:    "600 kiloohms",
	},
	{
		description: "blue and black and green",
		input:       []string{"blue", "black", "green"},
		expected:    "6 megaohms",
	},
	{
		description: "blue and black and blue",
		input:       []string{"blue", "black", "blue"},
		expected:    "60 megaohms",
	},
	{
		description: "blue and black and violet",
		input:       []string{"blue", "black", "violet"},
		expected:    "600 megaohms",
	},
	{
		description: "blue and black and grey",
		input:       []string{"blue", "black", "grey"},
		expected:    "6 gigaohms",
	},
	{
		description: "blue and black and white",
		input:       []string{"blue", "black", "white"},
		expected:    "60 gigaohms",
	},
	{
		description: "blue and brown and black",
		input:       []string{"blue", "brown", "black"},
		expected:    "61 ohms",
	},
	{
		description: "blue and brown and brown",
		input:       []string{"blue", "brown", "brown"},
		expected:    "610 ohms",
	},
	{
		description: "blue and brown and red",
		input:       []string{"blue", "brown", "red"},
		expected:    "6100 ohms",
	},
	{
		description: "blue and brown and orange",
		input:       []string{"blue", "brown", "orange"},
		expected:    "61 kiloohms",
	},
	{
		description: "blue and brown and yellow",
		input:       []string{"blue", "brown", "yellow"},
		expected:    "610 kiloohms",
	},
	{
		description: "blue and brown and green",
		input:       []string{"blue", "brown", "green"},
		expected:    "6100 kiloohms",
	},
	{
		description: "blue and brown and blue",
		input:       []string{"blue", "brown", "blue"},
		expected:    "61 megaohms",
	},
	{
		description: "blue and brown and violet",
		input:       []string{"blue", "brown", "violet"},
		expected:    "610 megaohms",
	},
	{
		description: "blue and brown and grey",
		input:       []string{"blue", "brown", "grey"},
		expected:    "6100 megaohms",
	},
	{
		description: "blue and brown and white",
		input:       []string{"blue", "brown", "white"},
		expected:    "61 gigaohms",
	},
	{
		description: "blue and red and black",
		input:       []string{"blue", "red", "black"},
		expected:    "62 ohms",
	},
	{
		description: "blue and red and brown",
		input:       []string{"blue", "red", "brown"},
		expected:    "620 ohms",
	},
	{
		description: "blue and red and red",
		input:       []string{"blue", "red", "red"},
		expected:    "6200 ohms",
	},
	{
		description: "blue and red and orange",
		input:       []string{"blue", "red", "orange"},
		expected:    "62 kiloohms",
	},
	{
		description: "blue and red and yellow",
		input:       []string{"blue", "red", "yellow"},
		expected:    "620 kiloohms",
	},
	{
		description: "blue and red and green",
		input:       []string{"blue", "red", "green"},
		expected:    "6200 kiloohms",
	},
	{
		description: "blue and red and blue",
		input:       []string{"blue", "red", "blue"},
		expected:    "62 megaohms",
	},
	{
		description: "blue and red and violet",
		input:       []string{"blue", "red", "violet"},
		expected:    "620 megaohms",
	},
	{
		description: "blue and red and grey",
		input:       []string{"blue", "red", "grey"},
		expected:    "6200 megaohms",
	},
	{
		description: "blue and red and white",
		input:       []string{"blue", "red", "white"},
		expected:    "62 gigaohms",
	},
	{
		description: "blue and orange and black",
		input:       []string{"blue", "orange", "black"},
		expected:    "63 ohms",
	},
	{
		description: "blue and orange and brown",
		input:       []string{"blue", "orange", "brown"},
		expected:    "630 ohms",
	},
	{
		description: "blue and orange and red",
		input:       []string{"blue", "orange", "red"},
		expected:    "6300 ohms",
	},
	{
		description: "blue and orange and orange",
		input:       []string{"blue", "orange", "orange"},
		expected:    "63 kiloohms",
	},
	{
		description: "blue and orange and yellow",
		input:       []string{"blue", "orange", "yellow"},
		expected:    "630 kiloohms",
	},
	{
		description: "blue and orange and green",
		input:       []string{"blue", "orange", "green"},
		expected:    "6300 kiloohms",
	},
	{
		description: "blue and orange and blue",
		input:       []string{"blue", "orange", "blue"},
		expected:    "63 megaohms",
	},
	{
		description: "blue and orange and violet",
		input:       []string{"blue", "orange", "violet"},
		expected:    "630 megaohms",
	},
	{
		description: "blue and orange and grey",
		input:       []string{"blue", "orange", "grey"},
		expected:    "6300 megaohms",
	},
	{
		description: "blue and orange and white",
		input:       []string{"blue", "orange", "white"},
		expected:    "63 gigaohms",
	},
	{
		description: "blue and yellow and black",
		input:       []string{"blue", "yellow", "black"},
		expected:    "64 ohms",
	},
	{
		description: "blue and yellow and brown",
		input:       []string{"blue", "yellow", "brown"},
		expected:    "640 ohms",
	},
	{
		description: "blue and yellow and red",
		input:       []string{"blue", "yellow", "red"},
		expected:    "6400 ohms",
	},
	{
		description: "blue and yellow and orange",
		input:       []string{"blue", "yellow", "orange"},
		expected:    "64 kiloohms",
	},
	{
		description: "blue and yellow and yellow",
		input:       []string{"blue", "yellow", "yellow"},
		expected:    "640 kiloohms",
	},
	{
		description: "blue and yellow and green",
		input:       []string{"blue", "yellow", "green"},
		expected:    "6400 kiloohms",
	},
	{
		description: "blue and yellow and blue",
		input:       []string{"blue", "yellow", "blue"},
		expected:    "64 megaohms",
	},
	{
		description: "blue and yellow and violet",
		input:       []string{"blue", "yellow", "violet"},
		expected:    "640 megaohms",
	},
	{
		description: "blue and yellow and grey",
		input:       []string{"blue", "yellow", "grey"},
		expected:    "6400 megaohms",
	},
	{
		description: "blue and yellow and white",
		input:       []string{"blue", "yellow", "white"},
		expected:    "64 gigaohms",
	},
	{
		description: "blue and green and black",
		input:       []string{"blue", "green", "black"},
		expected:    "65 ohms",
	},
	{
		description: "blue and green and brown",
		input:       []string{"blue", "green", "brown"},
		expected:    "650 ohms",
	},
	{
		description: "blue and green and red",
		input:       []string{"blue", "green", "red"},
		expected:    "6500 ohms",
	},
	{
		description: "blue and green and orange",
		input:       []string{"blue", "green", "orange"},
		expected:    "65 kiloohms",
	},
	{
		description: "blue and green and yellow",
		input:       []string{"blue", "green", "yellow"},
		expected:    "650 kiloohms",
	},
	{
		description: "blue and green and green",
		input:       []string{"blue", "green", "green"},
		expected:    "6500 kiloohms",
	},
	{
		description: "blue and green and blue",
		input:       []string{"blue", "green", "blue"},
		expected:    "65 megaohms",
	},
	{
		description: "blue and green and violet",
		input:       []string{"blue", "green", "violet"},
		expected:    "650 megaohms",
	},
	{
		description: "blue and green and grey",
		input:       []string{"blue", "green", "grey"},
		expected:    "6500 megaohms",
	},
	{
		description: "blue and green and white",
		input:       []string{"blue", "green", "white"},
		expected:    "65 gigaohms",
	},
	{
		description: "blue and blue and black",
		input:       []string{"blue", "blue", "black"},
		expected:    "66 ohms",
	},
	{
		description: "blue and blue and brown",
		input:       []string{"blue", "blue", "brown"},
		expected:    "660 ohms",
	},
	{
		description: "blue and blue and red",
		input:       []string{"blue", "blue", "red"},
		expected:    "6600 ohms",
	},
	{
		description: "blue and blue and orange",
		input:       []string{"blue", "blue", "orange"},
		expected:    "66 kiloohms",
	},
	{
		description: "blue and blue and yellow",
		input:       []string{"blue", "blue", "yellow"},
		expected:    "660 kiloohms",
	},
	{
		description: "blue and blue and green",
		input:       []string{"blue", "blue", "green"},
		expected:    "6600 kiloohms",
	},
	{
		description: "blue and blue and blue",
		input:       []string{"blue", "blue", "blue"},
		expected:    "66 megaohms",
	},
	{
		description: "blue and blue and violet",
		input:       []string{"blue", "blue", "violet"},
		expected:    "660 megaohms",
	},
	{
		description: "blue and blue and grey",
		input:       []string{"blue", "blue", "grey"},
		expected:    "6600 megaohms",
	},
	{
		description: "blue and blue and white",
		input:       []string{"blue", "blue", "white"},
		expected:    "66 gigaohms",
	},
	{
		description: "blue and violet and black",
		input:       []string{"blue", "violet", "black"},
		expected:    "67 ohms",
	},
	{
		description: "blue and violet and brown",
		input:       []string{"blue", "violet", "brown"},
		expected:    "670 ohms",
	},
	{
		description: "blue and violet and red",
		input:       []string{"blue", "violet", "red"},
		expected:    "6700 ohms",
	},
	{
		description: "blue and violet and orange",
		input:       []string{"blue", "violet", "orange"},
		expected:    "67 kiloohms",
	},
	{
		description: "blue and violet and yellow",
		input:       []string{"blue", "violet", "yellow"},
		expected:    "670 kiloohms",
	},
	{
		description: "blue and violet and green",
		input:       []string{"blue", "violet", "green"},
		expected:    "6700 kiloohms",
	},
	{
		description: "blue and violet and blue",
		input:       []string{"blue", "violet", "blue"},
		expected:    "67 megaohms",
	},
	{
		description: "blue and violet and violet",
		input:       []string{"blue", "violet", "violet"},
		expected:    "670 megaohms",
	},
	{
		description: "blue and violet and grey",
		input:       []string{"blue", "violet", "grey"},
		expected:    "6700 megaohms",
	},
	{
		description: "blue and violet and white",
		input:       []string{"blue", "violet", "white"},
		expected:    "67 gigaohms",
	},
	{
		description: "blue and grey and black",
		input:       []string{"blue", "grey", "black"},
		expected:    "68 ohms",
	},
	{
		description: "blue and grey and brown",
		input:       []string{"blue", "grey", "brown"},
		expected:    "680 ohms",
	},
	{
		description: "blue and grey and red",
		input:       []string{"blue", "grey", "red"},
		expected:    "6800 ohms",
	},
	{
		description: "blue and grey and orange",
		input:       []string{"blue", "grey", "orange"},
		expected:    "68 kiloohms",
	},
	{
		description: "blue and grey and yellow",
		input:       []string{"blue", "grey", "yellow"},
		expected:    "680 kiloohms",
	},
	{
		description: "blue and grey and green",
		input:       []string{"blue", "grey", "green"},
		expected:    "6800 kiloohms",
	},
	{
		description: "blue and grey and blue",
		input:       []string{"blue", "grey", "blue"},
		expected:    "68 megaohms",
	},
	{
		description: "blue and grey and violet",
		input:       []string{"blue", "grey", "violet"},
		expected:    "680 megaohms",
	},
	{
		description: "blue and grey and grey",
		input:       []string{"blue", "grey", "grey"},
		expected:    "6800 megaohms",
	},
	{
		description: "blue and grey and white",
		input:       []string{"blue", "grey", "white"},
		expected:    "68 gigaohms",
	},
	{
		description: "blue and white and black",
		input:       []string{"blue", "white", "black"},
		expected:    "69 ohms",
	},
	{
		description: "blue and white and brown",
		input:       []string{"blue", "white", "brown"},
		expected:    "690 ohms",
	},
	{
		description: "blue and white and red",
		input:       []string{"blue", "white", "red"},
		expected:    "6900 ohms",
	},
	{
		description: "blue and white and orange",
		input:       []string{"blue", "white", "orange"},
		expected:    "69 kiloohms",
	},
	{
		description: "blue and white and yellow",
		input:       []string{"blue", "white", "yellow"},
		expected:    "690 kiloohms",
	},
	{
		description: "blue and white and green",
		input:       []string{"blue", "white", "green"},
		expected:    "6900 kiloohms",
	},
	{
		description: "blue and white and blue",
		input:       []string{"blue", "white", "blue"},
		expected:    "69 megaohms",
	},
	{
		description: "blue and white and violet",
		input:       []string{"blue", "white", "violet"},
		expected:    "690 megaohms",
	},
	{
		description: "blue and white and grey",
		input:       []string{"blue", "white", "grey"},
		expected:    "6900 megaohms",
	},
	{
		description: "blue and white and white",
		input:       []string{"blue", "white", "white"},
		expected:    "69 gigaohms",
	},
	{
		description: "violet and black and black",
		input:       []string{"violet", "black", "black"},
		expected:    "70 ohms",
	},
	{
		description: "violet and black and brown",
		input:       []string{"violet", "black", "brown"},
		expected:    "700 ohms",
	},
	{
		description: "violet and black and red",
		input:       []string{"violet", "black", "red"},
		expected:    "7 kiloohms",
	},
	{
		description: "violet and black and orange",
		input:       []string{"violet", "black", "orange"},
		expected:    "70 kiloohms",
	},
	{
		description: "violet and black and yellow",
		input:       []string{"violet", "black", "yellow"},
		expected:    "700 kiloohms",
	},
	{
		description: "violet and black and green",
		input:       []string{"violet", "black", "green"},
		expected:    "7 megaohms",
	},
	{
		description: "violet and black and blue",
		input:       []string{"violet", "black", "blue"},
		expected:    "70 megaohms",
	},
	{
		description: "violet and black and violet",
		input:       []string{"violet", "black", "violet"},
		expected:    "700 megaohms",
	},
	{
		description: "violet and black and grey",
		input:       []string{"violet", "black", "grey"},
		expected:    "7 gigaohms",
	},
	{
		description: "violet and black and white",
		input:       []string{"violet", "black", "white"},
		expected:    "70 gigaohms",
	},
	{
		description: "violet and brown and black",
		input:       []string{"violet", "brown", "black"},
		expected:    "71 ohms",
	},
	{
		description: "violet and brown and brown",
		input:       []string{"violet", "brown", "brown"},
		expected:    "710 ohms",
	},
	{
		description: "violet and brown and red",
		input:       []string{"violet", "brown", "red"},
		expected:    "7100 ohms",
	},
	{
		description: "violet and brown and orange",
		input:       []string{"violet", "brown", "orange"},
		expected:    "71 kiloohms",
	},
	{
		description: "violet and brown and yellow",
		input:       []string{"violet", "brown", "yellow"},
		expected:    "710 kiloohms",
	},
	{
		description: "violet and brown and green",
		input:       []string{"violet", "brown", "green"},
		expected:    "7100 kiloohms",
	},
	{
		description: "violet and brown and blue",
		input:       []string{"violet", "brown", "blue"},
		expected:    "71 megaohms",
	},
	{
		description: "violet and brown and violet",
		input:       []string{"violet", "brown", "violet"},
		expected:    "710 megaohms",
	},
	{
		description: "violet and brown and grey",
		input:       []string{"violet", "brown", "grey"},
		expected:    "7100 megaohms",
	},
	{
		description: "violet and brown and white",
		input:       []string{"violet", "brown", "white"},
		expected:    "71 gigaohms",
	},
	{
		description: "violet and red and black",
		input:       []string{"violet", "red", "black"},
		expected:    "72 ohms",
	},
	{
		description: "violet and red and brown",
		input:       []string{"violet", "red", "brown"},
		expected:    "720 ohms",
	},
	{
		description: "violet and red and red",
		input:       []string{"violet", "red", "red"},
		expected:    "7200 ohms",
	},
	{
		description: "violet and red and orange",
		input:       []string{"violet", "red", "orange"},
		expected:    "72 kiloohms",
	},
	{
		description: "violet and red and yellow",
		input:       []string{"violet", "red", "yellow"},
		expected:    "720 kiloohms",
	},
	{
		description: "violet and red and green",
		input:       []string{"violet", "red", "green"},
		expected:    "7200 kiloohms",
	},
	{
		description: "violet and red and blue",
		input:       []string{"violet", "red", "blue"},
		expected:    "72 megaohms",
	},
	{
		description: "violet and red and violet",
		input:       []string{"violet", "red", "violet"},
		expected:    "720 megaohms",
	},
	{
		description: "violet and red and grey",
		input:       []string{"violet", "red", "grey"},
		expected:    "7200 megaohms",
	},
	{
		description: "violet and red and white",
		input:       []string{"violet", "red", "white"},
		expected:    "72 gigaohms",
	},
	{
		description: "violet and orange and black",
		input:       []string{"violet", "orange", "black"},
		expected:    "73 ohms",
	},
	{
		description: "violet and orange and brown",
		input:       []string{"violet", "orange", "brown"},
		expected:    "730 ohms",
	},
	{
		description: "violet and orange and red",
		input:       []string{"violet", "orange", "red"},
		expected:    "7300 ohms",
	},
	{
		description: "violet and orange and orange",
		input:       []string{"violet", "orange", "orange"},
		expected:    "73 kiloohms",
	},
	{
		description: "violet and orange and yellow",
		input:       []string{"violet", "orange", "yellow"},
		expected:    "730 kiloohms",
	},
	{
		description: "violet and orange and green",
		input:       []string{"violet", "orange", "green"},
		expected:    "7300 kiloohms",
	},
	{
		description: "violet and orange and blue",
		input:       []string{"violet", "orange", "blue"},
		expected:    "73 megaohms",
	},
	{
		description: "violet and orange and violet",
		input:       []string{"violet", "orange", "violet"},
		expected:    "730 megaohms",
	},
	{
		description: "violet and orange and grey",
		input:       []string{"violet", "orange", "grey"},
		expected:    "7300 megaohms",
	},
	{
		description: "violet and orange and white",
		input:       []string{"violet", "orange", "white"},
		expected:    "73 gigaohms",
	},
	{
		description: "violet and yellow and black",
		input:       []string{"violet", "yellow", "black"},
		expected:    "74 ohms",
	},
	{
		description: "violet and yellow and brown",
		input:       []string{"violet", "yellow", "brown"},
		expected:    "740 ohms",
	},
	{
		description: "violet and yellow and red",
		input:       []string{"violet", "yellow", "red"},
		expected:    "7400 ohms",
	},
	{
		description: "violet and yellow and orange",
		input:       []string{"violet", "yellow", "orange"},
		expected:    "74 kiloohms",
	},
	{
		description: "violet and yellow and yellow",
		input:       []string{"violet", "yellow", "yellow"},
		expected:    "740 kiloohms",
	},
	{
		description: "violet and yellow and green",
		input:       []string{"violet", "yellow", "green"},
		expected:    "7400 kiloohms",
	},
	{
		description: "violet and yellow and blue",
		input:       []string{"violet", "yellow", "blue"},
		expected:    "74 megaohms",
	},
	{
		description: "violet and yellow and violet",
		input:       []string{"violet", "yellow", "violet"},
		expected:    "740 megaohms",
	},
	{
		description: "violet and yellow and grey",
		input:       []string{"violet", "yellow", "grey"},
		expected:    "7400 megaohms",
	},
	{
		description: "violet and yellow and white",
		input:       []string{"violet", "yellow", "white"},
		expected:    "74 gigaohms",
	},
	{
		description: "violet and green and black",
		input:       []string{"violet", "green", "black"},
		expected:    "75 ohms",
	},
	{
		description: "violet and green and brown",
		input:       []string{"violet", "green", "brown"},
		expected:    "750 ohms",
	},
	{
		description: "violet and green and red",
		input:       []string{"violet", "green", "red"},
		expected:    "7500 ohms",
	},
	{
		description: "violet and green and orange",
		input:       []string{"violet", "green", "orange"},
		expected:    "75 kiloohms",
	},
	{
		description: "violet and green and yellow",
		input:       []string{"violet", "green", "yellow"},
		expected:    "750 kiloohms",
	},
	{
		description: "violet and green and green",
		input:       []string{"violet", "green", "green"},
		expected:    "7500 kiloohms",
	},
	{
		description: "violet and green and blue",
		input:       []string{"violet", "green", "blue"},
		expected:    "75 megaohms",
	},
	{
		description: "violet and green and violet",
		input:       []string{"violet", "green", "violet"},
		expected:    "750 megaohms",
	},
	{
		description: "violet and green and grey",
		input:       []string{"violet", "green", "grey"},
		expected:    "7500 megaohms",
	},
	{
		description: "violet and green and white",
		input:       []string{"violet", "green", "white"},
		expected:    "75 gigaohms",
	},
	{
		description: "violet and blue and black",
		input:       []string{"violet", "blue", "black"},
		expected:    "76 ohms",
	},
	{
		description: "violet and blue and brown",
		input:       []string{"violet", "blue", "brown"},
		expected:    "760 ohms",
	},
	{
		description: "violet and blue and red",
		input:       []string{"violet", "blue", "red"},
		expected:    "7600 ohms",
	},
	{
		description: "violet and blue and orange",
		input:       []string{"violet", "blue", "orange"},
		expected:    "76 kiloohms",
	},
	{
		description: "violet and blue and yellow",
		input:       []string{"violet", "blue", "yellow"},
		expected:    "760 kiloohms",
	},
	{
		description: "violet and blue and green",
		input:       []string{"violet", "blue", "green"},
		expected:    "7600 kiloohms",
	},
	{
		description: "violet and blue and blue",
		input:       []string{"violet", "blue", "blue"},
		expected:    "76 megaohms",
	},
	{
		description: "violet and blue and violet",
		input:       []string{"violet", "blue", "violet"},
		expected:    "760 megaohms",
	},
	{
		description: "violet and blue and grey",
		input:       []string{"violet", "blue", "grey"},
		expected:    "7600 megaohms",
	},
	{
		description: "violet and blue and white",
		input:       []string{"violet", "blue", "white"},
		expected:    "76 gigaohms",
	},
	{
		description: "violet and violet and black",
		input:       []string{"violet", "violet", "black"},
		expected:    "77 ohms",
	},
	{
		description: "violet and violet and brown",
		input:       []string{"violet", "violet", "brown"},
		expected:    "770 ohms",
	},
	{
		description: "violet and violet and red",
		input:       []string{"violet", "violet", "red"},
		expected:    "7700 ohms",
	},
	{
		description: "violet and violet and orange",
		input:       []string{"violet", "violet", "orange"},
		expected:    "77 kiloohms",
	},
	{
		description: "violet and violet and yellow",
		input:       []string{"violet", "violet", "yellow"},
		expected:    "770 kiloohms",
	},
	{
		description: "violet and violet and green",
		input:       []string{"violet", "violet", "green"},
		expected:    "7700 kiloohms",
	},
	{
		description: "violet and violet and blue",
		input:       []string{"violet", "violet", "blue"},
		expected:    "77 megaohms",
	},
	{
		description: "violet and violet and violet",
		input:       []string{"violet", "violet", "violet"},
		expected:    "770 megaohms",
	},
	{
		description: "violet and violet and grey",
		input:       []string{"violet", "violet", "grey"},
		expected:    "7700 megaohms",
	},
	{
		description: "violet and violet and white",
		input:       []string{"violet", "violet", "white"},
		expected:    "77 gigaohms",
	},
	{
		description: "violet and grey and black",
		input:       []string{"violet", "grey", "black"},
		expected:    "78 ohms",
	},
	{
		description: "violet and grey and brown",
		input:       []string{"violet", "grey", "brown"},
		expected:    "780 ohms",
	},
	{
		description: "violet and grey and red",
		input:       []string{"violet", "grey", "red"},
		expected:    "7800 ohms",
	},
	{
		description: "violet and grey and orange",
		input:       []string{"violet", "grey", "orange"},
		expected:    "78 kiloohms",
	},
	{
		description: "violet and grey and yellow",
		input:       []string{"violet", "grey", "yellow"},
		expected:    "780 kiloohms",
	},
	{
		description: "violet and grey and green",
		input:       []string{"violet", "grey", "green"},
		expected:    "7800 kiloohms",
	},
	{
		description: "violet and grey and blue",
		input:       []string{"violet", "grey", "blue"},
		expected:    "78 megaohms",
	},
	{
		description: "violet and grey and violet",
		input:       []string{"violet", "grey", "violet"},
		expected:    "780 megaohms",
	},
	{
		description: "violet and grey and grey",
		input:       []string{"violet", "grey", "grey"},
		expected:    "7800 megaohms",
	},
	{
		description: "violet and grey and white",
		input:       []string{"violet", "grey", "white"},
		expected:    "78 gigaohms",
	},
	{
		description: "violet and white and black",
		input:       []string{"violet", "white", "black"},
		expected:    "79 ohms",
	},
	{
		description: "violet and white and brown",
		input:       []string{"violet", "white", "brown"},
		expected:    "790 ohms",
	},
	{
		description: "violet and white and red",
		input:       []string{"violet", "white", "red"},
		expected:    "7900 ohms",
	},
	{
		description: "violet and white and orange",
		input:       []string{"violet", "white", "orange"},
		expected:    "79 kiloohms",
	},
	{
		description: "violet and white and yellow",
		input:       []string{"violet", "white", "yellow"},
		expected:    "790 kiloohms",
	},
	{
		description: "violet and white and green",
		input:       []string{"violet", "white", "green"},
		expected:    "7900 kiloohms",
	},
	{
		description: "violet and white and blue",
		input:       []string{"violet", "white", "blue"},
		expected:    "79 megaohms",
	},
	{
		description: "violet and white and violet",
		input:       []string{"violet", "white", "violet"},
		expected:    "790 megaohms",
	},
	{
		description: "violet and white and grey",
		input:       []string{"violet", "white", "grey"},
		expected:    "7900 megaohms",
	},
	{
		description: "violet and white and white",
		input:       []string{"violet", "white", "white"},
		expected:    "79 gigaohms",
	},
	{
		description: "grey and black and black",
		input:       []string{"grey", "black", "black"},
		expected:    "80 ohms",
	},
	{
		description: "grey and black and brown",
		input:       []string{"grey", "black", "brown"},
		expected:    "800 ohms",
	},
	{
		description: "grey and black and red",
		input:       []string{"grey", "black", "red"},
		expected:    "8 kiloohms",
	},
	{
		description: "grey and black and orange",
		input:       []string{"grey", "black", "orange"},
		expected:    "80 kiloohms",
	},
	{
		description: "grey and black and yellow",
		input:       []string{"grey", "black", "yellow"},
		expected:    "800 kiloohms",
	},
	{
		description: "grey and black and green",
		input:       []string{"grey", "black", "green"},
		expected:    "8 megaohms",
	},
	{
		description: "grey and black and blue",
		input:       []string{"grey", "black", "blue"},
		expected:    "80 megaohms",
	},
	{
		description: "grey and black and violet",
		input:       []string{"grey", "black", "violet"},
		expected:    "800 megaohms",
	},
	{
		description: "grey and black and grey",
		input:       []string{"grey", "black", "grey"},
		expected:    "8 gigaohms",
	},
	{
		description: "grey and black and white",
		input:       []string{"grey", "black", "white"},
		expected:    "80 gigaohms",
	},
	{
		description: "grey and brown and black",
		input:       []string{"grey", "brown", "black"},
		expected:    "81 ohms",
	},
	{
		description: "grey and brown and brown",
		input:       []string{"grey", "brown", "brown"},
		expected:    "810 ohms",
	},
	{
		description: "grey and brown and red",
		input:       []string{"grey", "brown", "red"},
		expected:    "8100 ohms",
	},
	{
		description: "grey and brown and orange",
		input:       []string{"grey", "brown", "orange"},
		expected:    "81 kiloohms",
	},
	{
		description: "grey and brown and yellow",
		input:       []string{"grey", "brown", "yellow"},
		expected:    "810 kiloohms",
	},
	{
		description: "grey and brown and green",
		input:       []string{"grey", "brown", "green"},
		expected:    "8100 kiloohms",
	},
	{
		description: "grey and brown and blue",
		input:       []string{"grey", "brown", "blue"},
		expected:    "81 megaohms",
	},
	{
		description: "grey and brown and violet",
		input:       []string{"grey", "brown", "violet"},
		expected:    "810 megaohms",
	},
	{
		description: "grey and brown and grey",
		input:       []string{"grey", "brown", "grey"},
		expected:    "8100 megaohms",
	},
	{
		description: "grey and brown and white",
		input:       []string{"grey", "brown", "white"},
		expected:    "81 gigaohms",
	},
	{
		description: "grey and red and black",
		input:       []string{"grey", "red", "black"},
		expected:    "82 ohms",
	},
	{
		description: "grey and red and brown",
		input:       []string{"grey", "red", "brown"},
		expected:    "820 ohms",
	},
	{
		description: "grey and red and red",
		input:       []string{"grey", "red", "red"},
		expected:    "8200 ohms",
	},
	{
		description: "grey and red and orange",
		input:       []string{"grey", "red", "orange"},
		expected:    "82 kiloohms",
	},
	{
		description: "grey and red and yellow",
		input:       []string{"grey", "red", "yellow"},
		expected:    "820 kiloohms",
	},
	{
		description: "grey and red and green",
		input:       []string{"grey", "red", "green"},
		expected:    "8200 kiloohms",
	},
	{
		description: "grey and red and blue",
		input:       []string{"grey", "red", "blue"},
		expected:    "82 megaohms",
	},
	{
		description: "grey and red and violet",
		input:       []string{"grey", "red", "violet"},
		expected:    "820 megaohms",
	},
	{
		description: "grey and red and grey",
		input:       []string{"grey", "red", "grey"},
		expected:    "8200 megaohms",
	},
	{
		description: "grey and red and white",
		input:       []string{"grey", "red", "white"},
		expected:    "82 gigaohms",
	},
	{
		description: "grey and orange and black",
		input:       []string{"grey", "orange", "black"},
		expected:    "83 ohms",
	},
	{
		description: "grey and orange and brown",
		input:       []string{"grey", "orange", "brown"},
		expected:    "830 ohms",
	},
	{
		description: "grey and orange and red",
		input:       []string{"grey", "orange", "red"},
		expected:    "8300 ohms",
	},
	{
		description: "grey and orange and orange",
		input:       []string{"grey", "orange", "orange"},
		expected:    "83 kiloohms",
	},
	{
		description: "grey and orange and yellow",
		input:       []string{"grey", "orange", "yellow"},
		expected:    "830 kiloohms",
	},
	{
		description: "grey and orange and green",
		input:       []string{"grey", "orange", "green"},
		expected:    "8300 kiloohms",
	},
	{
		description: "grey and orange and blue",
		input:       []string{"grey", "orange", "blue"},
		expected:    "83 megaohms",
	},
	{
		description: "grey and orange and violet",
		input:       []string{"grey", "orange", "violet"},
		expected:    "830 megaohms",
	},
	{
		description: "grey and orange and grey",
		input:       []string{"grey", "orange", "grey"},
		expected:    "8300 megaohms",
	},
	{
		description: "grey and orange and white",
		input:       []string{"grey", "orange", "white"},
		expected:    "83 gigaohms",
	},
	{
		description: "grey and yellow and black",
		input:       []string{"grey", "yellow", "black"},
		expected:    "84 ohms",
	},
	{
		description: "grey and yellow and brown",
		input:       []string{"grey", "yellow", "brown"},
		expected:    "840 ohms",
	},
	{
		description: "grey and yellow and red",
		input:       []string{"grey", "yellow", "red"},
		expected:    "8400 ohms",
	},
	{
		description: "grey and yellow and orange",
		input:       []string{"grey", "yellow", "orange"},
		expected:    "84 kiloohms",
	},
	{
		description: "grey and yellow and yellow",
		input:       []string{"grey", "yellow", "yellow"},
		expected:    "840 kiloohms",
	},
	{
		description: "grey and yellow and green",
		input:       []string{"grey", "yellow", "green"},
		expected:    "8400 kiloohms",
	},
	{
		description: "grey and yellow and blue",
		input:       []string{"grey", "yellow", "blue"},
		expected:    "84 megaohms",
	},
	{
		description: "grey and yellow and violet",
		input:       []string{"grey", "yellow", "violet"},
		expected:    "840 megaohms",
	},
	{
		description: "grey and yellow and grey",
		input:       []string{"grey", "yellow", "grey"},
		expected:    "8400 megaohms",
	},
	{
		description: "grey and yellow and white",
		input:       []string{"grey", "yellow", "white"},
		expected:    "84 gigaohms",
	},
	{
		description: "grey and green and black",
		input:       []string{"grey", "green", "black"},
		expected:    "85 ohms",
	},
	{
		description: "grey and green and brown",
		input:       []string{"grey", "green", "brown"},
		expected:    "850 ohms",
	},
	{
		description: "grey and green and red",
		input:       []string{"grey", "green", "red"},
		expected:    "8500 ohms",
	},
	{
		description: "grey and green and orange",
		input:       []string{"grey", "green", "orange"},
		expected:    "85 kiloohms",
	},
	{
		description: "grey and green and yellow",
		input:       []string{"grey", "green", "yellow"},
		expected:    "850 kiloohms",
	},
	{
		description: "grey and green and green",
		input:       []string{"grey", "green", "green"},
		expected:    "8500 kiloohms",
	},
	{
		description: "grey and green and blue",
		input:       []string{"grey", "green", "blue"},
		expected:    "85 megaohms",
	},
	{
		description: "grey and green and violet",
		input:       []string{"grey", "green", "violet"},
		expected:    "850 megaohms",
	},
	{
		description: "grey and green and grey",
		input:       []string{"grey", "green", "grey"},
		expected:    "8500 megaohms",
	},
	{
		description: "grey and green and white",
		input:       []string{"grey", "green", "white"},
		expected:    "85 gigaohms",
	},
	{
		description: "grey and blue and black",
		input:       []string{"grey", "blue", "black"},
		expected:    "86 ohms",
	},
	{
		description: "grey and blue and brown",
		input:       []string{"grey", "blue", "brown"},
		expected:    "860 ohms",
	},
	{
		description: "grey and blue and red",
		input:       []string{"grey", "blue", "red"},
		expected:    "8600 ohms",
	},
	{
		description: "grey and blue and orange",
		input:       []string{"grey", "blue", "orange"},
		expected:    "86 kiloohms",
	},
	{
		description: "grey and blue and yellow",
		input:       []string{"grey", "blue", "yellow"},
		expected:    "860 kiloohms",
	},
	{
		description: "grey and blue and green",
		input:       []string{"grey", "blue", "green"},
		expected:    "8600 kiloohms",
	},
	{
		description: "grey and blue and blue",
		input:       []string{"grey", "blue", "blue"},
		expected:    "86 megaohms",
	},
	{
		description: "grey and blue and violet",
		input:       []string{"grey", "blue", "violet"},
		expected:    "860 megaohms",
	},
	{
		description: "grey and blue and grey",
		input:       []string{"grey", "blue", "grey"},
		expected:    "8600 megaohms",
	},
	{
		description: "grey and blue and white",
		input:       []string{"grey", "blue", "white"},
		expected:    "86 gigaohms",
	},
	{
		description: "grey and violet and black",
		input:       []string{"grey", "violet", "black"},
		expected:    "87 ohms",
	},
	{
		description: "grey and violet and brown",
		input:       []string{"grey", "violet", "brown"},
		expected:    "870 ohms",
	},
	{
		description: "grey and violet and red",
		input:       []string{"grey", "violet", "red"},
		expected:    "8700 ohms",
	},
	{
		description: "grey and violet and orange",
		input:       []string{"grey", "violet", "orange"},
		expected:    "87 kiloohms",
	},
	{
		description: "grey and violet and yellow",
		input:       []string{"grey", "violet", "yellow"},
		expected:    "870 kiloohms",
	},
	{
		description: "grey and violet and green",
		input:       []string{"grey", "violet", "green"},
		expected:    "8700 kiloohms",
	},
	{
		description: "grey and violet and blue",
		input:       []string{"grey", "violet", "blue"},
		expected:    "87 megaohms",
	},
	{
		description: "grey and violet and violet",
		input:       []string{"grey", "violet", "violet"},
		expected:    "870 megaohms",
	},
	{
		description: "grey and violet and grey",
		input:       []string{"grey", "violet", "grey"},
		expected:    "8700 megaohms",
	},
	{
		description: "grey and violet and white",
		input:       []string{"grey", "violet", "white"},
		expected:    "87 gigaohms",
	},
	{
		description: "grey and grey and black",
		input:       []string{"grey", "grey", "black"},
		expected:    "88 ohms",
	},
	{
		description: "grey and grey and brown",
		input:       []string{"grey", "grey", "brown"},
		expected:    "880 ohms",
	},
	{
		description: "grey and grey and red",
		input:       []string{"grey", "grey", "red"},
		expected:    "8800 ohms",
	},
	{
		description: "grey and grey and orange",
		input:       []string{"grey", "grey", "orange"},
		expected:    "88 kiloohms",
	},
	{
		description: "grey and grey and yellow",
		input:       []string{"grey", "grey", "yellow"},
		expected:    "880 kiloohms",
	},
	{
		description: "grey and grey and green",
		input:       []string{"grey", "grey", "green"},
		expected:    "8800 kiloohms",
	},
	{
		description: "grey and grey and blue",
		input:       []string{"grey", "grey", "blue"},
		expected:    "88 megaohms",
	},
	{
		description: "grey and grey and violet",
		input:       []string{"grey", "grey", "violet"},
		expected:    "880 megaohms",
	},
	{
		description: "grey and grey and grey",
		input:       []string{"grey", "grey", "grey"},
		expected:    "8800 megaohms",
	},
	{
		description: "grey and grey and white",
		input:       []string{"grey", "grey", "white"},
		expected:    "88 gigaohms",
	},
	{
		description: "grey and white and black",
		input:       []string{"grey", "white", "black"},
		expected:    "89 ohms",
	},
	{
		description: "grey and white and brown",
		input:       []string{"grey", "white", "brown"},
		expected:    "890 ohms",
	},
	{
		description: "grey and white and red",
		input:       []string{"grey", "white", "red"},
		expected:    "8900 ohms",
	},
	{
		description: "grey and white and orange",
		input:       []string{"grey", "white", "orange"},
		expected:    "89 kiloohms",
	},
	{
		description: "grey and white and yellow",
		input:       []string{"grey", "white", "yellow"},
		expected:    "890 kiloohms",
	},
	{
		description: "grey and white and green",
		input:       []string{"grey", "white", "green"},
		expected:    "8900 kiloohms",
	},
	{
		description: "grey and white and blue",
		input:       []string{"grey", "white", "blue"},
		expected:    "89 megaohms",
	},
	{
		description: "grey and white and violet",
		input:       []string{"grey", "white", "violet"},
		expected:    "890 megaohms",
	},
	{
		description: "grey and white and grey",
		input:       []string{"grey", "white", "grey"},
		expected:    "8900 megaohms",
	},
	{
		description: "grey and white and white",
		input:       []string{"grey", "white", "white"},
		expected:    "89 gigaohms",
	},
	{
		description: "white and black and black",
		input:       []string{"white", "black", "black"},
		expected:    "90 ohms",
	},
	{
		description: "white and black and brown",
		input:       []string{"white", "black", "brown"},
		expected:    "900 ohms",
	},
	{
		description: "white and black and red",
		input:       []string{"white", "black", "red"},
		expected:    "9 kiloohms",
	},
	{
		description: "white and black and orange",
		input:       []string{"white", "black", "orange"},
		expected:    "90 kiloohms",
	},
	{
		description: "white and black and yellow",
		input:       []string{"white", "black", "yellow"},
		expected:    "900 kiloohms",
	},
	{
		description: "white and black and green",
		input:       []string{"white", "black", "green"},
		expected:    "9 megaohms",
	},
	{
		description: "white and black and blue",
		input:       []string{"white", "black", "blue"},
		expected:    "90 megaohms",
	},
	{
		description: "white and black and violet",
		input:       []string{"white", "black", "violet"},
		expected:    "900 megaohms",
	},
	{
		description: "white and black and grey",
		input:       []string{"white", "black", "grey"},
		expected:    "9 gigaohms",
	},
	{
		description: "white and black and white",
		input:       []string{"white", "black", "white"},
		expected:    "90 gigaohms",
	},
	{
		description: "white and brown and black",
		input:       []string{"white", "brown", "black"},
		expected:    "91 ohms",
	},
	{
		description: "white and brown and brown",
		input:       []string{"white", "brown", "brown"},
		expected:    "910 ohms",
	},
	{
		description: "white and brown and red",
		input:       []string{"white", "brown", "red"},
		expected:    "9100 ohms",
	},
	{
		description: "white and brown and orange",
		input:       []string{"white", "brown", "orange"},
		expected:    "91 kiloohms",
	},
	{
		description: "white and brown and yellow",
		input:       []string{"white", "brown", "yellow"},
		expected:    "910 kiloohms",
	},
	{
		description: "white and brown and green",
		input:       []string{"white", "brown", "green"},
		expected:    "9100 kiloohms",
	},
	{
		description: "white and brown and blue",
		input:       []string{"white", "brown", "blue"},
		expected:    "91 megaohms",
	},
	{
		description: "white and brown and violet",
		input:       []string{"white", "brown", "violet"},
		expected:    "910 megaohms",
	},
	{
		description: "white and brown and grey",
		input:       []string{"white", "brown", "grey"},
		expected:    "9100 megaohms",
	},
	{
		description: "white and brown and white",
		input:       []string{"white", "brown", "white"},
		expected:    "91 gigaohms",
	},
	{
		description: "white and red and black",
		input:       []string{"white", "red", "black"},
		expected:    "92 ohms",
	},
	{
		description: "white and red and brown",
		input:       []string{"white", "red", "brown"},
		expected:    "920 ohms",
	},
	{
		description: "white and red and red",
		input:       []string{"white", "red", "red"},
		expected:    "9200 ohms",
	},
	{
		description: "white and red and orange",
		input:       []string{"white", "red", "orange"},
		expected:    "92 kiloohms",
	},
	{
		description: "white and red and yellow",
		input:       []string{"white", "red", "yellow"},
		expected:    "920 kiloohms",
	},
	{
		description: "white and red and green",
		input:       []string{"white", "red", "green"},
		expected:    "9200 kiloohms",
	},
	{
		description: "white and red and blue",
		input:       []string{"white", "red", "blue"},
		expected:    "92 megaohms",
	},
	{
		description: "white and red and violet",
		input:       []string{"white", "red", "violet"},
		expected:    "920 megaohms",
	},
	{
		description: "white and red and grey",
		input:       []string{"white", "red", "grey"},
		expected:    "9200 megaohms",
	},
	{
		description: "white and red and white",
		input:       []string{"white", "red", "white"},
		expected:    "92 gigaohms",
	},
	{
		description: "white and orange and black",
		input:       []string{"white", "orange", "black"},
		expected:    "93 ohms",
	},
	{
		description: "white and orange and brown",
		input:       []string{"white", "orange", "brown"},
		expected:    "930 ohms",
	},
	{
		description: "white and orange and red",
		input:       []string{"white", "orange", "red"},
		expected:    "9300 ohms",
	},
	{
		description: "white and orange and orange",
		input:       []string{"white", "orange", "orange"},
		expected:    "93 kiloohms",
	},
	{
		description: "white and orange and yellow",
		input:       []string{"white", "orange", "yellow"},
		expected:    "930 kiloohms",
	},
	{
		description: "white and orange and green",
		input:       []string{"white", "orange", "green"},
		expected:    "9300 kiloohms",
	},
	{
		description: "white and orange and blue",
		input:       []string{"white", "orange", "blue"},
		expected:    "93 megaohms",
	},
	{
		description: "white and orange and violet",
		input:       []string{"white", "orange", "violet"},
		expected:    "930 megaohms",
	},
	{
		description: "white and orange and grey",
		input:       []string{"white", "orange", "grey"},
		expected:    "9300 megaohms",
	},
	{
		description: "white and orange and white",
		input:       []string{"white", "orange", "white"},
		expected:    "93 gigaohms",
	},
	{
		description: "white and yellow and black",
		input:       []string{"white", "yellow", "black"},
		expected:    "94 ohms",
	},
	{
		description: "white and yellow and brown",
		input:       []string{"white", "yellow", "brown"},
		expected:    "940 ohms",
	},
	{
		description: "white and yellow and red",
		input:       []string{"white", "yellow", "red"},
		expected:    "9400 ohms",
	},
	{
		description: "white and yellow and orange",
		input:       []string{"white", "yellow", "orange"},
		expected:    "94 kiloohms",
	},
	{
		description: "white and yellow and yellow",
		input:       []string{"white", "yellow", "yellow"},
		expected:    "940 kiloohms",
	},
	{
		description: "white and yellow and green",
		input:       []string{"white", "yellow", "green"},
		expected:    "9400 kiloohms",
	},
	{
		description: "white and yellow and blue",
		input:       []string{"white", "yellow", "blue"},
		expected:    "94 megaohms",
	},
	{
		description: "white and yellow and violet",
		input:       []string{"white", "yellow", "violet"},
		expected:    "940 megaohms",
	},
	{
		description: "white and yellow and grey",
		input:       []string{"white", "yellow", "grey"},
		expected:    "9400 megaohms",
	},
	{
		description: "white and yellow and white",
		input:       []string{"white", "yellow", "white"},
		expected:    "94 gigaohms",
	},
	{
		description: "white and green and black",
		input:       []string{"white", "green", "black"},
		expected:    "95 ohms",
	},
	{
		description: "white and green and brown",
		input:       []string{"white", "green", "brown"},
		expected:    "950 ohms",
	},
	{
		description: "white and green and red",
		input:       []string{"white", "green", "red"},
		expected:    "9500 ohms",
	},
	{
		description: "white and green and orange",
		input:       []string{"white", "green", "orange"},
		expected:    "95 kiloohms",
	},
	{
		description: "white and green and yellow",
		input:       []string{"white", "green", "yellow"},
		expected:    "950 kiloohms",
	},
	{
		description: "white and green and green",
		input:       []string{"white", "green", "green"},
		expected:    "9500 kiloohms",
	},
	{
		description: "white and green and blue",
		input:       []string{"white", "green", "blue"},
		expected:    "95 megaohms",
	},
	{
		description: "white and green and violet",
		input:       []string{"white", "green", "violet"},
		expected:    "950 megaohms",
	},
	{
		description: "white and green and grey",
		input:       []string{"white", "green", "grey"},
		expected:    "9500 megaohms",
	},
	{
		description: "white and green and white",
		input:       []string{"white", "green", "white"},
		expected:    "95 gigaohms",
	},
	{
		description: "white and blue and black",
		input:       []string{"white", "blue", "black"},
		expected:    "96 ohms",
	},
	{
		description: "white and blue and brown",
		input:       []string{"white", "blue", "brown"},
		expected:    "960 ohms",
	},
	{
		description: "white and blue and red",
		input:       []string{"white", "blue", "red"},
		expected:    "9600 ohms",
	},
	{
		description: "white and blue and orange",
		input:       []string{"white", "blue", "orange"},
		expected:    "96 kiloohms",
	},
	{
		description: "white and blue and yellow",
		input:       []string{"white", "blue", "yellow"},
		expected:    "960 kiloohms",
	},
	{
		description: "white and blue and green",
		input:       []string{"white", "blue", "green"},
		expected:    "9600 kiloohms",
	},
	{
		description: "white and blue and blue",
		input:       []string{"white", "blue", "blue"},
		expected:    "96 megaohms",
	},
	{
		description: "white and blue and violet",
		input:       []string{"white", "blue", "violet"},
		expected:    "960 megaohms",
	},
	{
		description: "white and blue and grey",
		input:       []string{"white", "blue", "grey"},
		expected:    "9600 megaohms",
	},
	{
		description: "white and blue and white",
		input:       []string{"white", "blue", "white"},
		expected:    "96 gigaohms",
	},
	{
		description: "white and violet and black",
		input:       []string{"white", "violet", "black"},
		expected:    "97 ohms",
	},
	{
		description: "white and violet and brown",
		input:       []string{"white", "violet", "brown"},
		expected:    "970 ohms",
	},
	{
		description: "white and violet and red",
		input:       []string{"white", "violet", "red"},
		expected:    "9700 ohms",
	},
	{
		description: "white and violet and orange",
		input:       []string{"white", "violet", "orange"},
		expected:    "97 kiloohms",
	},
	{
		description: "white and violet and yellow",
		input:       []string{"white", "violet", "yellow"},
		expected:    "970 kiloohms",
	},
	{
		description: "white and violet and green",
		input:       []string{"white", "violet", "green"},
		expected:    "9700 kiloohms",
	},
	{
		description: "white and violet and blue",
		input:       []string{"white", "violet", "blue"},
		expected:    "97 megaohms",
	},
	{
		description: "white and violet and violet",
		input:       []string{"white", "violet", "violet"},
		expected:    "970 megaohms",
	},
	{
		description: "white and violet and grey",
		input:       []string{"white", "violet", "grey"},
		expected:    "9700 megaohms",
	},
	{
		description: "white and violet and white",
		input:       []string{"white", "violet", "white"},
		expected:    "97 gigaohms",
	},
	{
		description: "white and grey and black",
		input:       []string{"white", "grey", "black"},
		expected:    "98 ohms",
	},
	{
		description: "white and grey and brown",
		input:       []string{"white", "grey", "brown"},
		expected:    "980 ohms",
	},
	{
		description: "white and grey and red",
		input:       []string{"white", "grey", "red"},
		expected:    "9800 ohms",
	},
	{
		description: "white and grey and orange",
		input:       []string{"white", "grey", "orange"},
		expected:    "98 kiloohms",
	},
	{
		description: "white and grey and yellow",
		input:       []string{"white", "grey", "yellow"},
		expected:    "980 kiloohms",
	},
	{
		description: "white and grey and green",
		input:       []string{"white", "grey", "green"},
		expected:    "9800 kiloohms",
	},
	{
		description: "white and grey and blue",
		input:       []string{"white", "grey", "blue"},
		expected:    "98 megaohms",
	},
	{
		description: "white and grey and violet",
		input:       []string{"white", "grey", "violet"},
		expected:    "980 megaohms",
	},
	{
		description: "white and grey and grey",
		input:       []string{"white", "grey", "grey"},
		expected:    "9800 megaohms",
	},
	{
		description: "white and grey and white",
		input:       []string{"white", "grey", "white"},
		expected:    "98 gigaohms",
	},
	{
		description: "white and white and black",
		input:       []string{"white", "white", "black"},
		expected:    "99 ohms",
	},
	{
		description: "white and white and brown",
		input:       []string{"white", "white", "brown"},
		expected:    "990 ohms",
	},
	{
		description: "white and white and red",
		input:       []string{"white", "white", "red"},
		expected:    "9900 ohms",
	},
	{
		description: "white and white and orange",
		input:       []string{"white", "white", "orange"},
		expected:    "99 kiloohms",
	},
	{
		description: "white and white and yellow",
		input:       []string{"white", "white", "yellow"},
		expected:    "990 kiloohms",
	},
	{
		description: "white and white and green",
		input:       []string{"white", "white", "green"},
		expected:    "9900 kiloohms",
	},
	{
		description: "white and white and blue",
		input:       []string{"white", "white", "blue"},
		expected:    "99 megaohms",
	},
	{
		description: "white and white and violet",
		input:       []string{"white", "white", "violet"},
		expected:    "990 megaohms",
	},
	{
		description: "white and white and grey",
		input:       []string{"white", "white", "grey"},
		expected:    "9900 megaohms",
	},
	{
		description: "white and white and white",
		input:       []string{"white", "white", "white"},
		expected:    "99 gigaohms",
	},
}
