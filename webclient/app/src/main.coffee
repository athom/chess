getCanvas = () ->
  width = 1280.0
  height = 800.0

  c = document.getElementById("chess-board")
  ctx = c.getContext("2d")

  c.width  = window.innerWidth
  c.height = window.innerHeight
  ratiow = c.width/width
  ratioh = c.height/height
  if ratiow < ratioh
          ratio = ratiow
          c.height = height * ratio
  else
          ratio = ratioh
          c.width = width * ratio

  ctx.scale(ratio, ratio)

  return new Canvas(ctx, width, height)


main = () ->
  canvas = getCanvas()
  world = new World(canvas)

  status_bar = new StatusBar
  parser = new Parser(world, status_bar)
  ws = new Websocket(parser)
  ws.connect()

$ ->
  main()
