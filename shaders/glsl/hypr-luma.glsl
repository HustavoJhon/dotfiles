// PACMAN - Hyprland luminance effect shader
// Subtle luminance-based color enhancement

uniform float time;
uniform vec2 resolution;

vec4 main() {
    vec2 uv = gl_FragCoord.xy / resolution.xy;
    vec4 color = texture(iChannel0, uv);

    float luma = dot(color.rgb, vec3(0.299, 0.587, 0.114));
    float enhancement = 1.0 + 0.1 * (1.0 - luma);

    vec3 enhanced = color.rgb * enhancement;

    float vignette = 1.0 - 0.15 * length(uv - 0.5) * 1.4;
    enhanced *= vignette;

    return vec4(enhanced, color.a);
}
