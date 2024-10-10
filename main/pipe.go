components {
  id: "pipe"
  component: "/main/pipe.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"end_pipe_red\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/pipe.atlas\"\n"
  "}\n"
  ""
  position {
    z: 1.0
  }
}
