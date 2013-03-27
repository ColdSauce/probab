package stat

import (
	"fmt"
	"testing"
)

// Test against R:moments:agostino

func TestSampleAgostino(t *testing.T) {
	const delta = 1e-4

	d := []float64{-0.670789480, -0.150969186, -0.551925020, -0.482406038, -1.173156862, 0.367903204, 1.059533774, 1.794411979, -1.895192051, -0.571951213, -0.726043212, 0.627210778, -1.272177905, -0.144797609, -0.968422674, 0.283461379, 0.988415637, 2.475978661, -0.536749821, -0.025513855, -1.534900705, -0.230832466, 0.620245249, 1.495396229, -2.111838253, 0.228080184, 0.251295357, 0.181974246, -0.001878552, -2.085836898, 0.834083317, -1.321850451, 1.192865393, -0.195088801, 1.748176403, 1.780233827, -1.212874246, -0.275506872, -0.474797265, -0.496230943, 1.872632489, 0.165565088, -1.143077239, -0.683540167, -2.029626407, 0.027388799, 0.263192684, -0.637438762, -0.460420680, 0.316033043, -0.166534773, -0.270718891, 0.634184308, -0.302168613, 0.928756911, 0.737386257, -0.546999206, 0.107967686, -0.082850172, -1.895955213, 1.576817423, -0.079801919, 0.382725131, 0.758249530, -2.054212936, -1.285879785, -0.009835473, 0.636242575, 0.169251475, -0.465556922, -0.387162464, -0.172489252, -1.910986672, 2.291340848, 1.499807228, -0.707866084, -0.430782756, -0.943412247, 1.663281221, 0.392045957, -0.006054918, -0.330406096, 0.519286029, 0.700169631, -0.521821484, -0.412689406, 0.814297193, -0.709301954, -0.861132419, -1.471550480, -0.686849429, 0.174766897, -0.333436632, 1.291139538, 0.040605740, -0.399557291, 2.169674201, -1.712460224, 0.377688521, -1.831853999}
	fmt.Println("Agostino test two-sided")
	alt := 0
	a, b, c := Agostino(d, alt)
	x, y, z := 0.2013, 0.5719, 0.5674
	if abs(x-a) > delta {
		fmt.Println("failed skewness: ", a, x)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed z: ", b, y)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed p-value: ", c, z)
		t.Error()
	}

	fmt.Println("Agostino test less")
	alt = 1
	a, b, c = Agostino(d, alt)
	x, y, z = 0.2013, 0.5719, 0.2837
	if abs(x-a) > delta {
		fmt.Println("failed skewness: ", a, x)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed z: ", b, y)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed p-value: ", c, z)
		t.Error()
	}

	fmt.Println("Agostino test greater")
	alt = 2
	a, b, c = Agostino(d, alt)
	x, y, z = 0.2013, 0.5719, 0.2837
	if abs(x-a) > delta {
		fmt.Println("failed skewness: ", a, x)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed z: ", b, y)
		t.Error()
	}
	if abs(y-b) > delta {
		fmt.Println("failed p-value: ", c, z)
		t.Error()
	}
}
