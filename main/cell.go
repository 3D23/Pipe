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
