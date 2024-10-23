components {
  id: "pipe"
  component: "/main/pipe.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"end_pipe_red\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 115.0\n"
  "  y: 115.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/pipe.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.9
  }
  scale {
    z: 0.9
  }
}
