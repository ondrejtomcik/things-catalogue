package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type ThingModel struct {
	Context     []string            `json:"@context"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Properties  map[string]Property `json:"properties"`
	Actions     map[string]Action   `json:"actions"`
	Events      map[string]Event    `json:"events"`
}

type Property struct {
	Type        string              `json:"type"`
	Description string              `json:"description"`
	Properties  map[string]Property `json:"properties,omitempty"`
}

type Action struct {
	Description string `json:"description"`
}

type Event struct {
	Description string `json:"description"`
}

func propNames() []string {
	return []string{
		"temperature", "humidity", "status", "brightness", "volume", "battery", "altitude", "pressure", "speed", "capacity",
		"voltage", "current", "frequency", "phase", "powerFactor", "energyConsumption", "activePower", "reactivePower",
		"apparentPower", "chargeStatus", "chargeLevel", "remainingChargeTime", "connectorType", "gridStatus", "breakerStatus",
		"transformerLoad", "insulationResistance", "earthResistance", "arcFlashRisk", "oilLevel", "oilQuality", "sf6Pressure",
		"sf6Quality", "busbarLoad", "neutralLoad", "tripStatus", "faultIndicator", "circuitIntegrity", "groundingStatus",
		"surgeProtectorStatus", "overloadIndicator", "harmonics", "waveformDistortion", "varistorStatus", "capacitorHealth",
		"inductorHealth", "relayStatus", "meterReading", "tariffRate", "operationalMode", "emergencyMode", "protectionSetpoint",
		"backupBatteryStatus", "circuitIdentifier", "feedDirection", "loadShedding", "demandResponse", "gridTieStatus",
		"inverterEfficiency", "renewableSourceInput", "solarPanelHealth", "windTurbineEfficiency", "thermalSensor", "vibrationSensor",
		"communicationStatus", "dataRate", "latency", "uptime", "firmwareVersion", "maintenanceStatus", "operationalHours",
		"totalTrips", "manualOverride", "lastInspectionDate", "nextMaintenanceDate", "assetID", "locationIdentifier",
		"remoteControlCapability", "synchronizationStatus", "backupGeneratorStatus", "fuelLevel", "environmentalImpact",
		"carbonEmission", "totalYield", "efficiencyRating", "safetyRating", "faultCode", "diagnosticMessage", "operationalLog",
	}
}

func propTypes() []string {
	return []string{"string", "number", "boolean", "object"}
}

func randomChoice(items []string, remove bool) (string, []string) {
	if len(items) == 0 {
		return "", items
	}
	index := rand.Intn(len(items))
	chosen := items[index]

	if remove {
		items = append(items[:index], items[index+1:]...)
	}

	return chosen, items
}

func randomString(prefix string) string {
	return fmt.Sprintf("%s%x", prefix, rand.Int31())
}

type ManufacturerDetails struct {
	Name     string
	Families []string
	Products []string
}

func randomProperties() map[string]Property {
	propNames := propNames()
	propTypes := propTypes()
	properties := make(map[string]Property)
	propCount := rand.Intn(5) + 1

	for i := 0; i < propCount; i++ {
		var propName string
		propName, propNames = randomChoice(propNames, true) // Remove chosen name
		if propName == "" {
			break // Exit the loop if we run out of property names
		}

		propType, _ := randomChoice(propTypes, false) // Do not remove chosen type

		if propType == "object" {
			nestedProps := randomNestedProperties()
			if len(nestedProps) > 0 {
				properties[propName] = Property{
					Type:        "object",
					Description: fmt.Sprintf("An object containing nested properties for %s", propName),
					Properties:  nestedProps,
				}
				continue
			}
		}

		properties[propName] = Property{
			Type:        propType,
			Description: fmt.Sprintf("The %s of the device", propName),
		}
	}

	return properties
}

func generateManufacturersData() []ManufacturerDetails {
	var data []ManufacturerDetails

	for i := 0; i < 5; i++ { // Generate 5 manufacturers
		manufacturer := ManufacturerDetails{Name: randomString("Mfg")}

		for j := 0; j < 10; j++ { // 10 families per manufacturer
			manufacturer.Families = append(manufacturer.Families, randomString("Family"))
		}

		for k := 0; k < 20; k++ { // 20 products per family
			manufacturer.Products = append(manufacturer.Products, randomString("Product"))
		}

		data = append(data, manufacturer)
	}

	return data
}

func mutateProperties(props map[string]Property) map[string]Property {
	// Modify a property type or add a new one
	propNames := propNames()
	propTypes := propTypes()

	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}

	// Change a random property's type
	if len(keys) > 0 {
		randKey := keys[rand.Intn(len(keys))]
		property := props[randKey] // Extract the property
		newType, _ := randomChoice(propTypes, false)
		property.Type = newType // Modify the extracted property

		if newType == "object" && property.Properties == nil {
			property.Properties = randomNestedProperties()
		}

		props[randKey] = property // Put the modified property back into the map
	}

	// Add a new property
	unusedNames := make([]string, 0)
	for _, name := range propNames {
		if _, exists := props[name]; !exists {
			unusedNames = append(unusedNames, name)
		}
	}

	if len(unusedNames) > 0 {
		newPropName, _ := randomChoice(unusedNames, true)
		newType, _ := randomChoice(propTypes, false)
		if newType == "object" {
			props[newPropName] = Property{
				Type:        "object",
				Description: fmt.Sprintf("An object containing nested properties for %s", newPropName),
				Properties:  randomNestedProperties(),
			}
		} else {
			props[newPropName] = Property{
				Type:        newType,
				Description: fmt.Sprintf("The %s of the device", newPropName),
			}
		}
	}

	return props
}

func randomNestedProperties() map[string]Property {
	propNames := propNames()
	propTypes := propTypes()

	properties := make(map[string]Property)
	propCount := rand.Intn(3) + 1

	for i := 0; i < propCount; i++ {
		propName, updatedPropNames := randomChoice(propNames, true)
		propNames = updatedPropNames // Update the slice with the remaining names

		propType, _ := randomChoice(propTypes, false) // We don't remove types, so ignoring the second return value

		properties[propName] = Property{
			Type:        propType,
			Description: fmt.Sprintf("The %s of the nested object", propName),
		}
	}

	return properties
}

var manufacturersData = generateManufacturersData()

func main() {
	rand.Seed(time.Now().UnixNano())

	for _, manufacturer := range manufacturersData {
		for _, family := range manufacturer.Families {
			for _, product := range manufacturer.Products {
				model := ThingModel{
					Context:     []string{"https://www.w3.org/2019/wot/td/v1", "https://www.w3.org/2019/wot/tm"},
					Title:       product,
					Description: fmt.Sprintf("%s by %s", product, manufacturer.Name),
					Properties:  randomProperties(),
					Actions:     map[string]Action{"toggle": {Description: "Toggle the device state"}},
					Events:      map[string]Event{"alert": {Description: "An alert event"}},
				}

				// Create directory structure: Manufacturer/Family/Product
				dir := fmt.Sprintf("./%s/%s/%s", manufacturer.Name, family, product)
				os.MkdirAll(dir, os.ModePerm)

				// Generate 1 to 3 versions of thing models for each product
				model.Properties = randomProperties() // Initial set of properties for version 1
				versions := rand.Intn(3) + 1
				for v := 1; v <= versions; v++ {
					filename := fmt.Sprintf("%s/%s.v%d.jsonld", dir, product, v)
					data, _ := json.MarshalIndent(model, "", "  ")
					ioutil.WriteFile(filename, data, os.ModePerm)
					if v < versions {
						model.Properties = mutateProperties(model.Properties) // Change properties for next version
					}
				}
			}
		}
	}
}
