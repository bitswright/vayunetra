#include "wifi_manager.h"

#define WIFI_SSID       "WIFI_SSID"
#define WIFI_PASSWORD   "WIFI_PASWRD"

#define WIFI_TIMEOUT_MS 15000 // 15s

void wifiConnect() {
    Serial.print("Connecting to WiFi");

    WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

    unsigned long startAttempt = millis();

    while (Wifi.status() != WL_CONNECTED) {
        delay(500);
        Serial.print(".");

        if (millis() - startAttempt > WIFI_TIMEOUT_MS) {
            Serial.println("\nWiFi connection Failed (timeout)");
            WiFi.disconnect();
            return false;
        }
    }

    Serial.println("\nWiFi connected");
    Serial.print("IP address: ");
    Serial.println(WiFi.localIP());
    return true;
}