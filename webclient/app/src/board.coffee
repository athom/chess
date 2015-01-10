class Board
  WIDTH = 600
  HEIGHT = 600
  COLOR_BORDER = 'RGB(100, 0, 240)'
  COLOR_LINE = 'RGB(100, 0, 240)'

  constructor: (size) ->
    @size = size

  radius: ->
    WIDTH/@size/2.0 - 3

  position2point: (pos) ->
    if pos.x < 0 or pos.x >= @size or pos.y < 0 or pos.y >= @size
      return {x: -1, y: -1}

    widthUnit = WIDTH/@size
    heightUnit = HEIGHT/@size
    x = -WIDTH/2 + pos.x*widthUnit + widthUnit/2
    y = -HEIGHT/2 + pos.y*heightUnit + heightUnit/2
    return {x: x, y: -y}


  render: (canvas) ->
    # boarder rectangle
    canvas.drawRect(COLOR_BORDER, {
      x:-WIDTH/2,
      y:-HEIGHT/2,
      w:WIDTH,
      h:HEIGHT
    })

    # inside lines
    widthUnit = WIDTH/@size
    heightUnit = HEIGHT/@size
    for i in [1...@size]
      i = i-(@size)/2
      x1 = i*widthUnit
      y1 = i*heightUnit
      canvas.drawLine(COLOR_BORDER, x1, -HEIGHT/2, x1, HEIGHT/2)
      canvas.drawLine(COLOR_BORDER, -HEIGHT/2, y1, HEIGHT/2, y1)
