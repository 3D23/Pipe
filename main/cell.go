components {
  id: "cell"
  component: "/main/cell.script"
}
embedded_components {
  id: "cell_sprite"
  type: "sprite"
  data: "default_animation: \"gridCell\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 64.0\n"
  "  y: 64.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/cell.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "cleat_factory"
  type: "factory"
  data: "prototype: \"/main/cleat.go\"\n"
  "load_dynamically: true\n"
  ""
}
embedded_components {
  id: "pipe_factory"
  type: "factory"
  data: "prototype: \"/main/pipe.go\"\n"
  ""
}
