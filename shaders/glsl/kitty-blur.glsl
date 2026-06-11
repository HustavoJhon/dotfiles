// PACMAN.DOTS - Kitty background blur shader
// Gaussian blur approximation for transparent backgrounds

const vec3  sampleBrighten = vec3(0.4, 0.4, 0.4);
const float sampleStrength = 0.6;
const float sampleScale = 0.003;

vec4 main() {
    vec2 uv = gl_FragCoord.xy / iResolution.xy;
    vec4 color = texture(iChannel0, uv);

    vec2 samples[9] = vec2[](
        vec2(-1.0, -1.0), vec2(0.0, -1.0), vec2(1.0, -1.0),
        vec2(-1.0,  0.0), vec2(0.0,  0.0), vec2(1.0,  0.0),
        vec2(-1.0,  1.0), vec2(0.0,  1.0), vec2(1.0,  1.0)
    );

    vec3 blurred = vec3(0.0);
    float weights[9] = float[](
        0.0625, 0.125, 0.0625,
        0.125,  0.25,  0.125,
        0.0625, 0.125, 0.0625
    );

    for (int i = 0; i < 9; i++) {
        vec2 offset = samples[i] * sampleScale;
        blurred += texture(iChannel0, uv + offset).rgb * weights[i];
    }

    color.rgb = mix(color.rgb, blurred * sampleBrighten, sampleStrength);
    return color;
}
