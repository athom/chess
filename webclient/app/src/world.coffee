class World
  constructor: (canvas) ->
    @canvas = canvas
    @objects = []
    @board = new Board(6)
    @units = []

  init_units: (board_info) ->
    @units = []
    for unit in board_info.Units
      @units.push new Unit(
        @board,
        unit.side,
        unit.value,
        {
          x: unit.pos.X,
          y: unit.pos.Y
        }
      )

  render: () ->
    @canvas.clear()
    @board.render(@canvas)
    for unit in @units
      unit.render(@canvas)
