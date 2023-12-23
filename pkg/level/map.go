package level

// Create a Map list of strings
// Each string is a row of the map
// Each character of the string is representing a tile
// 11 rows of 28 tiles each
var Map = [][]string{
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", "X", "X", " "},
	{" ", "X", "X", "X", "X", " ", " ", "P", " ", " ", " ", " ", "X", "X", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
	{" ", "X", "X", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", "X", " ", " ", "X", "X", "X", "X", " ", " ", " ", " ", "X", "X", " ", " ", "X", "X", "X", " ", " ", " "},
	{" ", " ", " ", " ", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
	{"X", "X", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", "X", "X", "X", "X", " ", " ", "X", "X", " ", " ", "X", "X", "X", "X", " ", " "},
}