package main

import "log"

// kita bikin aturan disini

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

// make profile ini buat bikin profile. returnnya Profile struct
func MakeProfile(name string) Profile {

	if name == "Goku" {
		return Profile{
			Name:   name,
			Health: 200,
			Power:  100,
			Exp:    0,
		}
	}
	if name == "Sasuke" {
		return Profile{
			Name:   name,
			Health: 400,
			Power:  100,
			Exp:    0,
		}
	}
	if name == "Naruto" {
		return Profile{
			Name:   name,
			Health: 400,
			Power:  100,
			Exp:    0,
		}
	}

	return Profile{}

}

func PowerUp(profile Profile, multiplier int) Profile {
	return Profile{

		Name:   profile.Name,
		Health: profile.Health * multiplier,
		Power:  profile.Power * multiplier,
		Exp:    profile.Exp * multiplier,
	}
}

func main() {

	log.Println("======Goku================================")

	var goku = MakeProfile("Goku")
	log.Println("Name : ", goku.Name)
	log.Println("Health : ", goku.Health)
	log.Println("Power :", goku.Power)
	log.Println("Exp : ", goku.Exp)
	log.Println("======PowerUp ================================")
	var PowerUpGoku = PowerUp(goku, 2)

	log.Println("Name : ", PowerUpGoku.Name)
	log.Println("Helth : ", PowerUpGoku.Health)
	log.Println("Power :", PowerUpGoku.Power)
	log.Println("Exp :", PowerUpGoku.Exp)

	log.Println("========Sasuke============")

	var sasuke = MakeProfile("Sasuke")
	log.Println("Name : ", sasuke.Name)
	log.Println("Health: ", sasuke.Health)
	log.Println("Power : ", sasuke.Power)
	log.Println("Exp: ", sasuke.Exp)

	log.Println("======PowerUp================================")

	var PowerUpSasuke = PowerUp(sasuke, 5)
	log.Println("Name : ", PowerUpSasuke.Name)
	log.Println("Health :", PowerUpSasuke.Health)
	log.Println("Power : ", PowerUpSasuke.Power)
	log.Println("Exp : ", PowerUpSasuke.Exp)

	log.Println("========Naruto============")

	naruto := MakeProfile("Naruto")
	log.Println("Name : ", naruto.Name)
	log.Println("Health: ", naruto.Health)
	log.Println("Power : ", naruto.Power)
	log.Println("Exp: ", naruto.Exp)

	log.Println("========PowerUp============")

	PowerUpNaruto := PowerUp(naruto, 5)
	log.Println("Name : ", PowerUpNaruto.Name)
	log.Println("Health :", PowerUpNaruto.Health)
	log.Println("Power : ", PowerUpNaruto.Power)
	log.Println("Exp : ", PowerUpNaruto.Exp)

}
