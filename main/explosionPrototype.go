components {
  id: "explosion"
  component: "/scriptsRepository/explosion.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"explosion\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/graphicAssets.atlas\"\n"
  "}\n"
  ""
}
