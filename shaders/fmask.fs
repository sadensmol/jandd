#version 330 core

in vec2 fragTexCoord;

uniform sampler2D texture0; // background
uniform sampler2D texture1; // mask

out vec4 finalColor;

void main() {

    vec4 c1 = texture(texture0, fragTexCoord);
    vec4 c2 = texture(texture1, fragTexCoord);

    if (c2.a == 1.0) { // if pixel isn't transparent then we draw background
        finalColor = c1;
        return;
    }


    discard;
}