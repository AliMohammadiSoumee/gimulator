file = open('initMoves.txt', 'w')

initMoves = set()

def addSquareWithUpLeftPoint(x, y):
	initMoves.add((i, j, i + 1, j))
	initMoves.add((i, j, i + 1, j + 1))
	initMoves.add((i, j, i, j + 1))
	initMoves.add((i + 1, j, i + 1, j + 1))
	initMoves.add((i + 1, j, i, j + 1))
	initMoves.add((i, j + 1, i + 1, j + 1))

for i in range(1, 12):
	for j in range(1, 14):
		if (i == 1 or i == 10) and j < 13:
			addSquareWithUpLeftPoint(i, j)
		elif (j == 1 or j == 12) and i != 5 and i != 6 and i < 10:
			addSquareWithUpLeftPoint(i, j)

initMoves.add((5, 1, 6, 1))
initMoves.add((6, 1, 7, 1))
initMoves.add((5, 13, 6, 13))
initMoves.add((6, 13, 7, 13))

for move in initMoves:
	x, y, a, b = move
	file.write("{{A: State{{X: {}, Y: {}}}, B: State{{X: {}, Y: {}}}}},\n".format(x, y, a, b))
