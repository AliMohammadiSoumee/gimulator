file = open('initMoves.txt', 'w')

initMoves = set()

def addSquareWithUpLeftPoint(x, y):
	initMoves.add((i, j, i + 1, j))
	initMoves.add((i, j, i + 1, j + 1))
	initMoves.add((i, j, i, j + 1))
	initMoves.add((i + 1, j, i + 1, j + 1))
	initMoves.add((i + 1, j, i, j + 1))
	initMoves.add((i, j + 1, i + 1, j + 1))

for i in range(1, 13):
	for j in range(1, 9):
		if (i == 1 or i == 12) and j != 5 and j != 6:
			addSquareWithUpLeftPoint(i, j)
		elif j == 1 or j == 10:
			addSquareWithUpLeftPoint(i, j)

initMoves.add((1, 5, 1, 6))
initMoves.add((1, 6, 1, 7))
initMoves.add((13, 5, 1, 6))
initMoves.add((13, 6, 13, 7))

for move in initMoves:
	x, y, a, b = move
	file.write("{{A: types.State{{X: {}, Y: {}}}, B: types.State{{X: {}, Y: {}}}}},\n".format(x, y, a, b))
