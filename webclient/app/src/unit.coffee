class Unit
  COLOR_CIRCLE = 'RGB(244, 122, 2)'
  COLOR_UNIT_BLACK = 'RGB(0, 0, 0)'
  COLOR_UNIT_WHITE = 'RGB(255, 255, 255)'

  constructor: (board, side, value, pos) ->
    @board = board
    @side = side
    @value = value
    @pos = pos

  render: (canvas) ->
    point = @board.position2point @pos
    r = @board.radius()
    if @side == 1
      canvas.drawCircle(COLOR_CIRCLE, point.x, point.y, r)
      canvas.fillCircle(COLOR_UNIT_BLACK, point.x, point.y, r)
      canvas.drawText(COLOR_UNIT_WHITE, "80px Arial", @value, point.x, point.y+30)
    else
      canvas.drawCircle(COLOR_CIRCLE, point.x, point.y, r)
      canvas.drawText(COLOR_UNIT_BLACK, "80px Arial", @value, point.x, point.y+30)
