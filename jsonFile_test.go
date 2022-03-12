package main
import "testing"
func TestJsonFile(t *testing.T){
	message := "We are 2F-Capital Intership Go developers!"
	testerText := jsonMarshaling(message)
	achievedMSG := jsonMarshaling(message)
	requiredMSG := testerText

	if achievedMSG != requiredMSG {
		t.Errorf("Achieved Message is %q Required Message is %q", achievedMSG, requiredMSG)
	}
}