class World
  constructor: () ->
    @objects = []
    @ws = new Websocket
    @ws.connect()


  register: (obj) ->
    @objects.push obj

  render: (canvas) ->
    for obj in @objects
      obj.render(canvas)
