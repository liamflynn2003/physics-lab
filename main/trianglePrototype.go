components {
  id: "triangleScript"
  component: "/scriptsRepository/triangleScript.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"triangle1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/graphicAssets.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/main/triangle.convexshape\"\n"
  "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  ""
}
