components {
  id: "cleat"
  component: "/main/cleat.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"cleat_red\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/cleat.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
