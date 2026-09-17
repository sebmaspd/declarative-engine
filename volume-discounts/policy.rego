package volumediscounts

import rego.v1

# Section X(a): Base Pricing
base_unit_price := 50.00

# Section X(b): Discount Tiers, keyed off cumulative purchase volume.
default discount := 0

discount := 0.05 if {
	input.purchase_volume >= 500
	input.purchase_volume < 1500
}

discount := 0.10 if {
	input.purchase_volume >= 1500
	input.purchase_volume < 5000
}

discount := 0.15 if {
	input.purchase_volume >= 5000
}

default tier := "None"

tier := "Silver" if {
	input.purchase_volume >= 500
	input.purchase_volume < 1500
}

tier := "Gold" if {
	input.purchase_volume >= 1500
	input.purchase_volume < 5000
}

tier := "Platinum" if {
	input.purchase_volume >= 5000
}

net_unit_price := base_unit_price * (1 - discount)

volume_discounted_price := net_unit_price * input.purchase_volume

result := {
	"tier": tier,
	"unit_price": base_unit_price,
	"discount": discount,
	"net_unit_price": net_unit_price,
	"volume_discounted_price": volume_discounted_price,
}
