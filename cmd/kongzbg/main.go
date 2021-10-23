package main

import (
	"log"
	"fmt"
	"os"
	"image"
	"encoding/json"
	"io/ioutil"
	// "image/gif"
	"image/png"
	"image/color"
)

const kongzPixelInRealPixels = 20




func main() {
	log.Printf("Hello!")

	processKongz(6)
}

func processKongz(id uint) {
	log.Printf("Processing kongz %d", id)

    encodedImageFile, err := os.Open(fmt.Sprintf("assets/images/%d", id))
    if err != nil {
        log.Panic(err)
    }
    defer encodedImageFile.Close()

    imageData, _, err := image.Decode(encodedImageFile)
    if err != nil {
        log.Panic(err)
    }

    // Copy the image in an editable data struct
	newImageData := image.NewRGBA(imageData.Bounds())
	size := newImageData.Bounds().Size()
	for x := 0; x < size.X; x++ {
	    for y := 0; y < size.Y; y++ {
	        newImageData.Set(x, y, imageData.At(x, y))
	    }
	}

	// Set the background as transparent
    for kongzY, maskLine := range getKongzBackgroundMask(id) {
    	for kongzX, mask := range maskLine {
    		if mask == 1 {
				for x := kongzX * kongzPixelInRealPixels; x < (kongzX + 1) * kongzPixelInRealPixels; x++ {
				    for y := kongzY * kongzPixelInRealPixels; y < (kongzY + 1) * kongzPixelInRealPixels; y++ {
				        newImageData.Set(x, y, color.RGBA{0, 0, 0, 0})
				    }
				}    			
    		}
    	}
    }

    // Save the result
    outfile, err := os.Create(fmt.Sprintf("assets/images-nobg/%d", id))
    if err != nil {
        log.Panic(err)
    }
    defer outfile.Close()
    png.Encode(outfile, newImageData)
}

func getKongzBackgroundMask(id uint) (mask [][]int) {


	// Common mask to all
	var mainBackgroundMask = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
	}

	// Init the mask
	mask = mainBackgroundMask

	// Get the genesis traits
	kongzTraits := getKongzTraitByType("Genesis Trait", id)
	log.Printf("Traits: %v", kongzTraits)

	// Substractive mask per trait
	var traitsSubstractiveMasks = []struct{
		traits []string
		maskLocationsToUnset [][]int // X/Y format
	}{
		{
			traits: []string{"Propeller Hat"},
			maskLocationsToUnset: [][]int{{5, 10}, {5, 9}, {5, 8}, {6, 8}, {6, 7}, {7, 7}, {7, 6}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {13, 4}, {13, 3}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {21, 3}, {22, 3}, {23, 4}, {24, 4}, {25, 5}, {26, 7}, {26, 6}, {27, 8}, {27, 7}, {28, 8}, {28, 9}, {17, 2}, {17, 1}, {16, 1}, {15, 1}, {14, 1}, {13, 1}, {12, 1}, {11, 1}, {18, 1}, {19, 1}, {20, 1}, {21, 1}, {22, 1}, {23, 1}},
		},
		{
			traits: []string{"Blue Sunglasses", "Red Sunglasses"},
			maskLocationsToUnset: [][]int{{29, 13}, {29, 14}},
		},
		{
			traits: []string{"Headphones"},
			maskLocationsToUnset: [][]int{{9, 6}, {10, 6}, {10, 5}, {10, 4}, {11, 5}, {11, 4}, {11, 3}, {12, 4}, {12, 3}, {13, 4}, {13, 3}, {13, 2}, {14, 4}, {14, 3}, {14, 2}, {15, 4}, {15, 3}, {15, 2}, {16, 3}, {16, 2}, {17, 3}, {17, 2}, {18, 3}, {18, 2}, {19, 3}, {19, 2}, {20, 3}, {20, 2}, {21, 3}},
		},
		{
			traits: []string{"Flat Top"},
			maskLocationsToUnset: [][]int{{4, 12}, {5, 10}, {6, 8}, {7, 7}, {7, 6}, {8, 7}, {8, 6}, {8, 5}, {9, 6}, {9, 5}, {9, 4}, {10, 6}, {10, 5}, {10, 4}, {10, 3}, {11, 5}, {11, 4}, {11, 3}, {11, 2}, {12, 4}, {12, 3}, {12, 2}, {13, 4}, {13, 3}, {13, 2}, {13, 1}, {14, 4}, {14, 3}, {14, 2}, {14, 1}, {15, 4}, {15, 3}, {15, 2}, {15, 1}, {16, 3}, {16, 2}, {16, 1}, {17, 3}, {17, 2}, {17, 1}, {18, 3}, {18, 2}, {18, 1}, {19, 3}, {19, 2}, {19, 1}, {20, 3}, {20, 2}, {20, 1}, {21, 3}, {21, 2}, {21, 1}, {22, 3}, {22, 2}, {22, 1}, {23, 4}, {23, 3}, {23, 2}, {23, 1}, {24, 4}, {24, 3}, {24, 2}, {24, 1}, {25, 5}, {25, 4}, {25, 3}, {25, 2}, {25, 1}, {26, 7}, {26, 6}, {26, 5}, {26, 4}, {26, 3}, {26, 2}, {26, 1}, {27, 8}, {27, 7}, {27, 6}, {27, 5}, {27, 4}, {28, 9}, {28, 8}, {28, 7}},
		},
		{
			traits: []string{"Purple Bandana"},
			maskLocationsToUnset: [][]int{{2, 11}, {3, 11}, {3, 10}, {4, 10}, {5, 10}, {5, 9}, {5, 8}, {4, 8}, {3, 8}, {2, 8}, {2, 7}, {3, 7}, {6, 8}, {7, 7}, {8, 7}, {9, 6}, {10, 6}, {10, 5}, {11, 5}, {11, 4}, {12, 4}, {13, 4}, {13, 3}, {14, 4}, {14, 3}, {15, 4}, {15, 3}, {16, 3}, {17, 3}, {18, 3}, {19, 3}, {20, 3}, {21, 3}, {22, 3}, {23, 4}, {24, 4}, {25, 5}, {26, 6}, {26, 7}, {27, 8}},
		},
	}

	// Apply the mask
	for _, kongzTrait := range kongzTraits {
		for _, traitsSubstractiveMask := range traitsSubstractiveMasks {
			for _, traitOfSubstractiveMask := range traitsSubstractiveMask.traits {
				if kongzTrait == traitOfSubstractiveMask {
					// Ok our kongz needs to get this substractive mask
					for _, maskLocationToUnset := range traitsSubstractiveMask.maskLocationsToUnset {
						mask[maskLocationToUnset[1]][maskLocationToUnset[0]] = 0
					}
				}
			}
		}
	}

	return
}


type Attribute struct {
	TraitType string `json:"trait_type"`
	Value string `json:"value"`
}

type Metadata struct {
	Image string `json:"image"`
	ExternalUrl string `json:"external_url"`
	Name string `json:"name"`
	Attributes []Attribute `json:"attributes"`
}

func getKongzTraitByType(traitType string, id uint) (traits []string) {
	// Read metadata
	jsonFile, err := os.Open(fmt.Sprintf("assets/metadata/%d", id))
	if err != nil {
	    fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var metadata Metadata
	json.Unmarshal(byteValue, &metadata)

	for _, attribute := range metadata.Attributes {
		if attribute.TraitType == traitType {
			traits = append(traits, attribute.Value)
		}
	}

	return
}