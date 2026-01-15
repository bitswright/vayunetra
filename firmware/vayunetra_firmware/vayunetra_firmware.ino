#include "wifi_manager.h"

void setup() {
    Serial.begin(115200);
    delay(1000);

    Serial.println("Booting VayuNetra...");
    bool connected = wifiConnect();

    if(!connected) {
        Serial.println("Running in offline mode");
    }
}

void loop() {
    // nothing here yet
}