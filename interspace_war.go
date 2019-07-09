package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)


type OFFICER struct {
	name string
	age int
}

type Pilot struct {

	EvasionSkill int
}

type Engineer struct {
	DefenceSkill int
}

type Captain struct {
	TradingSkill int
}

type WeaponMaster struct {
	defenseSkill int
	AttackSkill int
}

type Miner struct {
	MiningSkill int
}

type TravelShip struct {
	name string
	age int
	evasionskill int
    Defenseskill int
    tradingskill int
    defenseskill int
    attackskill int
    miningskill int
	ArmourStrength int
	speed int
	Fuel int
	Gems int
	VALOR int
	trophy int
}
var T TravelShip
type Resource struct {
	Fuel int
	Gems int
}

type Planet struct {
planet string
alien string
}

type alienspecies struct {
	alien string
	attackskill int
	defenseskill int
}

type AlienSpaceship struct {
	alien string
	ArmourStrength int
	evasionskill int
	speed int
	attackskill int
}

func main(){
var t int
var pil Pilot
var Eng Engineer
var Cap Captain
var WM WeaponMaster
var Min Miner
T.trophy=1
fmt.Print("Welcome To The Game\n")
fmt.Print("\n")
time.Sleep(2*time.Second)
fmt.Print("Before Starting The Game, please select the travelship->\n")
time.Sleep(2*time.Second)
fmt.Print("1: RED TravelShip---->\n"  +
	             "PILOT->\n"  +
	             "\t\tEVASION SKILL: 10\n" +
	             "ENGINEER->\n"  +
	             "\t\tDEFENCE SKILL: 5\n"  +
	             "CAPTAIN->\n"  +
	             "\t\tTRADING SKILL: 5\n"  +
	             "WEAPON MASTER->\n"  +
	             "\t\tDEFENSE SKILL: 10\n\t\tATTACK SKILL: 11\n"  +
	             "MINER->\n"  +
	             "\t\tMINING SKILL: 5\n")
	time.Sleep(2*time.Second)
	fmt.Printf("2: BLUE TravelShip---->\n"  +
	             "PILOT->\n"  +
	             "\t\tEVASION SKILL: 9\n" +
	             "ENGINEER->\n"  +
	             "\t\tDEFENCE SKILL: 4\n"  +
	             "CAPTAIN->\n"  +
	             "\t\tTRADING SKILL: 5\n"  +
	             "WEAPON MASTER->\n"  +
	             "\t\tDEFENSE SKILL: 10\n\t\tATTACK SKILL: 11\n"  +
	             "MINER->\n"  +
	             "\t\tMINING SKILL: 5\n")
	time.Sleep(2*time.Second)
	fmt.Printf("3: GREEN TravelShip---->\n"  +
	             "PILOT->\n"  +
	             "\t\tEVASION SKILL: 10\n"  +
	             "ENGINEER->\n"  +
	             "\t\tDEFENCE SKILL: 6\n"  +
	             "CAPTAIN->\n"  +
	             "\t\tTRADING SKILL: 5\n"  +
	             "WEAPON MASTER->\n"  +
	             "\t\tDEFENSE SKILL: 11\n\t\tATTACK SKILL: 10\n"  +
	             "MINER->\n"  +
	             "\t\tMINING SKILL: 5\n")
	time.Sleep(2*time.Second)
	fmt.Printf("4: WHITE TravelShip---->\n"  +
	             "PILOT->\n"  +
	             "\t\tEVASION SKILL: 10\n"  +
	             "ENGINEER->\n"  +
	             "\t\tDEFENCE SKILL: 5\n"  +
	             "CAPTAIN:->\n" +
	             "\t\tTRADING SKILL: 5\n"  +
	             "WEAPON MASTER->\n"  +
	             "\t\tDEFENSE SKILL: 12\n\t\tATTACK SKILL: 9\n"  +
	             "MINER->\n"  +
	             "\t\tMINING SKILL: 5\n")
	time.Sleep(2*time.Second)
	fmt.Printf("5: BLACK Travelship---->\n"  +
	             "PILOT->\n"  +
	             "\t\tEVASION SKILL: 10\n"  +
	             "ENGINEER->\n"  +
	             "\t\tDEFENCE SKILL: 4\n"  +
	             "CAPTAIN->\n"  +
	             "\t\tTRADING SKILL: 5\n"  +
	             "WEAPON MASTER->\n"  +
	             "\t\tDEFENSE SKILL: 9\n\t\tATTACK SKILL: 12\n"  +
	             "MINER->\n"  +
	             "\t\tMINING SKILL: 5\n")
	time.Sleep(2*time.Second)
    fmt.Printf("\tSelect(1-5): ")
	fmt.Scan(&t)
	switch t {
	case 1:
		fmt.Printf(" \tSo you have chosen RED TravelShip\n")
		pil.EvasionSkill = 10
		Eng.DefenceSkill = 5
		Cap.TradingSkill = 5
		WM.defenseSkill = 10
		WM.AttackSkill = 11
		Min.MiningSkill = 5
		assigning(T, pil, Eng, Cap, WM, Min)
	case 2:
		fmt.Printf(" \tSo you have chosen BLUE TravelShip\n")
		pil.EvasionSkill = 9
		Eng.DefenceSkill = 4
		Cap.TradingSkill = 5
		WM.defenseSkill = 10
		WM.AttackSkill = 11
		Min.MiningSkill = 5
		assigning(T, pil, Eng, Cap, WM, Min)
	case 3:
		fmt.Printf(" \tSo you have chosen GREEN TravelShip\n")
		pil.EvasionSkill = 10
		Eng.DefenceSkill = 6
		Cap.TradingSkill = 5
		WM.defenseSkill = 11
		WM.AttackSkill = 10
		Min.MiningSkill = 5
		assigning(T, pil, Eng, Cap, WM, Min)
	case 4:
		fmt.Printf(" \tSo you have chosen WHITE TravelShip\n")
		pil.EvasionSkill = 10
		Eng.DefenceSkill = 5
		Cap.TradingSkill = 5
		WM.defenseSkill = 12
		WM.AttackSkill = 9
		Min.MiningSkill = 5
		assigning(T, pil, Eng, Cap, WM, Min)
	case 5:
		fmt.Printf(" \tSo you have chosen BLACK TravelShip\n")
		pil.EvasionSkill = 10
		Eng.DefenceSkill = 4
		Cap.TradingSkill = 5
		WM.defenseSkill = 9
		WM.AttackSkill = 12
		Min.MiningSkill = 5
		assigning(T, pil, Eng, Cap, WM, Min)
	default:
		fmt.Printf("\tTHEIR IS NO SHIP FOR THIS OPTION\n" +
			              "\tCHOOSE OTHER OPTION\n")
		time.Sleep(2*time.Second)
	}
}

func assigning(TS TravelShip,pil Pilot,Eng Engineer, Cap Captain,WM WeaponMaster,Min Miner) {
var o int
var OFFC OFFICER
var Res Resource
time.Sleep(2*time.Second)
fmt.Printf("\n\t\tGOOD MORNING OFFICER\n")
time.Sleep(2*time.Second)
fmt.Printf("\tBEFORE WE START THE JOURNEY, MAY I HAVE YOUR NAME,PLEASE: ")
fmt.Scan(&OFFC.name)
time.Sleep(2*time.Second)
fmt.Print("\tWE WOULD ASO LIKE TO KNOW YOUR AGE,PLEASE: ")
fmt.Scan(&OFFC.age)
fmt.Printf("\n")
TS.name=OFFC.name
TS.age=OFFC.age
TS.evasionskill=pil.EvasionSkill
TS.Defenseskill=Eng.DefenceSkill
TS.tradingskill=Cap.TradingSkill
TS.defenseskill=WM.defenseSkill
TS.attackskill=WM.AttackSkill
TS.miningskill=Min.MiningSkill
Res.Fuel=200
Res.Gems=200
TS.Fuel=Res.Fuel
TS.Gems=Res.Gems
TS.ArmourStrength=10
TS.speed=10
T=TS
fmt.Print("\n\n\t\tSO, LETS START THE JOURNEY\n")
time.Sleep(2*time.Second)
fmt.Print("\nSTORY->\n\t Space: Its the final frontier, the place where no one can hear you scream, and a \n" +
		"\tboundless backdrop that squashes many man's ego. Man is developing and going far and far into the\n" +
		"\tspace to discover things beyond his imagination.\n\n" +
		"\tIt has been several centuries since man went for the first time in space. Current year is 2600.\n" +
		"\tWe are not alone and not the most powerful in this vast universe. War has been going on between\n" +
		"\tmultiple planets. And now we are not adventuring into the space but fighting for survival. So \n" +
		"\tlets see, who survives and who is the conqueror of space.\n\n")
time.Sleep(4*time.Second)
fmt.Printf(	"\n\tLET THE BATTLE BEGIN")
time.Sleep(2*time.Second)
fmt.Printf("\n\n")
i:=1
for i<1000 {
	if T.VALOR>50{
		fmt.Printf("SINCE YOU HAVE %d VALOR POINT(S), YOU WON MASTER TROPHY %d",T.VALOR,T.trophy)
		T.trophy+=1
		T.VALOR=0
	}
	if T.Fuel<150 {
		time.Sleep(2*time.Second)
		fmt.Printf("\n\tYou have insufficient fuel, lets restore\n")
     fuelrestore()
	}
	if T.Gems<150{
		time.Sleep(2*time.Second)
		fmt.Printf("\n\tYou have insufficient gems, lets restore\n")
		Gemsrestore()
	}

	fmt.Printf("\n\t\t%s, we are starting the journey\n",TS.name )
	time.Sleep(2*time.Second)
	fmt.Printf("\t\tPlease choose one of the options\n\t" +
		              "\t\t1: Planet Search for resources\n\t" +
		              "\t\t2: AlienSpaceShip for battle\n\t" +
		              "\t\t3: Check Status\n\t" +
		              "\t\t4: Exit Game\n")
	fmt.Printf("\n\t\tYOUR OPTION: ")
	fmt.Scan(&o)
	switch o {
	    case 1:
                T.Fuel-=10
	    	    planetsearch(T)
	    case 2:
	    	    T.Fuel-=10
	    	    AlienSpaceShip(T)
	    case 3:
                check()
	    case 4:
	    	    os.Exit(1)
	    default:
	    	fmt.Print("\tChoose correct option\n")
			time.Sleep(2*time.Second)
	}

	i+=i
}

}
func check(){
	time.Sleep(2*time.Second)
	fmt.Printf("YOUR STATUS---->\n")
	fmt.Printf("\tNAME OF OFFICER-->%s\n",T.name)
	fmt.Printf("\tAGE OF OFFICER-->%d\n",T.age)
	fmt.Printf("\tEVASION SKILL OF PILOT-->%d\n",T.evasionskill)
	fmt.Printf("\tDEFENSE SKILL OF ENGINEER-->%d\n",T.Defenseskill)
	fmt.Printf("\tTRADING SKILL OF ENGINEER-->%d\n",T.tradingskill)
	fmt.Printf("\tATTACK SKILL OF WEAPON MASTER-->%d\n",T.attackskill)
	fmt.Printf("\tDEFENSE SKILL OF WEAPON MASTER-->%d\n",T.defenseskill)
	fmt.Printf("\tMINING SKILL OF MINER-->%d\n",T.miningskill)
	fmt.Printf("\tFUEL ON TRAVELSHIP-->%d\n",T.Fuel)
	fmt.Printf("\tGems ON TRAVELSHIP-->%d\n",T.Gems)
	fmt.Printf("\tSPEED OF TRAVELSHIP-->%d\n",T.speed)
	fmt.Printf("\t MASTER TROPHY--> %d\n",T.trophy)
	fmt.Printf("\t VALOR POINTS-->%d\n\n",T.VALOR)
}
func planetsearch(TS TravelShip) {
	var  f int
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tYOU ARE GOING ON AN EXPLORATION. DO YOU WANT TO " +
		"\n\t1.MINE?" +
		"\n\t2.TRADE?" +
		"\n\tSELECT 1 FOR TRADE OR 2 FOR MINE AND ANY OTHER NUMBER FOR GOING BACK TO START ADVENTURE SCREEN" +
		"\n\tSELECT( 1 OR 2): ")
	fmt.Scan(&f)
	if f == 1 || f==2{

		planetandaliens(TS,f)
		return

	} else {
		fmt.Printf("\n\tLETS GO BACK TO OUR PLANET")
		return
	}
}
func AlienSpaceShip(TS TravelShip){
var t int
var ASP AlienSpaceship
t=rand.Intn(13)
switch t{
case 0:
	ASP.alien="GALVATRON"
	battleonspaceship(TS,ASP)
	return
case 1:
	ASP.alien="NORAX"
	battleonspaceship(TS,ASP)
	return
case 2:
	ASP.alien="EQUINOX"
	battleonspaceship(TS,ASP)
	return
case 3:
	ASP.alien="METATRON"
	battleonspaceship(TS,ASP)
	return
case 4:
	ASP.alien="DALPHIN"
	battleonspaceship(TS,ASP)
	return
case 5:
	ASP.alien="NEOTIA"
	battleonspaceship(TS,ASP)
	return
case 6:
	ASP.alien="MENTROS"
	battleonspaceship(TS,ASP)
	return
case 7:
	ASP.alien="YUUNARIS"
	battleonspaceship(TS,ASP)
	return
case 8:
	ASP.alien="LESCO"
	battleonspaceship(TS,ASP)
	return
case 9:
	ASP.alien="REVEREND"
	battleonspaceship(TS,ASP)
	return
default:
	fmt.Printf("\n\tUNABLE TO FIND AN ALIENSPACESHIP\n")
	time.Sleep(2*time.Second)
	fmt.Printf("LETS GO BACK TO OUR PLANET")
	time.Sleep(2*time.Second)
	return
}
}
func battleonplanet(TS TravelShip,P Planet,AS alienspecies,f int){
	var pt,pa,r,m int
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tYOU HAVE LANDED ON PLANET '%s'",P.planet)
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tYou have encountered the alienspecies of the planet: %s",P.alien)
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tThey seems to be hostile towards your team.")
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tPREPARE FOR BATTLE")
	time.Sleep(2*time.Second)
	for i := 0; i < 3; i++ {
		fmt.Printf("\n\n\tROUND %d-->", i+1)
		if TS.attackskill > AS.defenseskill {
			pt += 1
			TS.attackskill += 1
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tYOUR ATTACK WAS HIGHER THAN THEIR DEFENSE")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tYOU GAIN ONE POINT")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tYOUR BATTLESPIRIT INCREASED YOUR ATTACK BY ONE POINT TEMPORARILY")
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tATTACK---->%d\n",TS.attackskill)
			time.Sleep(2*time.Second)
		} else {
			pa += 1
			TS.attackskill += 1
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tYOUR ATTACK WAS LOWER THAN THEIR DEFENSE")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tTHE ALIENS GAIN ONE POINT")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tYOUR BATTLESPIRIT INCREASED YOUR ATTACK BY ONE POINT TEMPORARILY")
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tATTACK---->%d\n",TS.attackskill)
			time.Sleep(2*time.Second)
		}
		if TS.defenseskill > AS.attackskill {
			pt += 1
			TS.defenseskill += 1
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tYOUR DEFENSE WAS HIGHER THAN THEIR ATTACK")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tYOU GAIN ONE POINT")
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tYOUR BATTLESPIRIT INCREASED YOUR DEFENSE BY ONE POINT TEMPORARILY")
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tDEFENSE---->%d\n",TS.defenseskill)
			time.Sleep(2*time.Second)
		} else {
			pa += 1
			TS.defenseskill += 1
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tYOUR DEFENSE WAS LOWER THAN THEIR ATTACK")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tTHE ALIENS GAIN ONE POINT")
			time.Sleep(2*time.Second)
			fmt.Printf(	"\n\tYOUR BATTLESPIRIT INCREASED YOUR DEFENSE BY ONE POINT TEMPORARILY")
			time.Sleep(2*time.Second)
			fmt.Printf("\n\tDEFENSE---->%d\n",TS.defenseskill)
			time.Sleep(2*time.Second)
		}

	}
	if pt > pa {
		fmt.Printf("\n\tYOUR POINTS--->%d\n",pt)
		fmt.Printf("\n\tALIEN POINTS--->%d\n",pa)
		time.Sleep(2*time.Second)
		fmt.Printf("\n\t YOU HAVE %d POINT(S) HIGHER THAN THE ELORIS", pt-pa)
		time.Sleep(2*time.Second)
		fmt.Printf("\n\tYOU WON")
		time.Sleep(2*time.Second)
		r=1
		time.Sleep(2*time.Second)
	} else if pa >= pt {
		fmt.Printf("\n\tYOUR POINTS--->%d\n",pt)
		fmt.Printf("\n\tALIEN POINTS--->%d\n",pa)
		time.Sleep(2*time.Second)
		fmt.Printf("\nTHE ALIENSPECIES HAVE %d POINT(S) HIGHER THAN YOU", pa-pt)
		time.Sleep(2*time.Second)
		fmt.Printf("\nYOU LOSE")
		r=0
		time.Sleep(2*time.Second)
		fmt.Printf("\n\tLETS GO BACK TO OUR PLANET")
		time.Sleep(2*time.Second)
	}
	if r==1{
		if f==1{
			fmt.Printf("\n\tYOU ARE ELIGIBLE FOR MINING\n")
			time.Sleep(2*time.Second)
			m=1
			mining(P,m)
		} else {
			fmt.Printf("\n\tYOU ARE ELIGIBLE FOR TRADING\n")
			time.Sleep(2*time.Second)
			m=1
			trade(P,m)
		}

	} else {
		return
		}
}
func planetandaliens(TS TravelShip,f int) {
	var P Planet
	var AS alienspecies
	t:= rand.Intn(10)
	var m int = 0
	c := rand.Intn(10)
	switch t {
	case 0:
		P.planet = "ALLANA"
		P.alien = "ELORIS"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(15)
		AS.defenseskill = rand.Intn(15)

		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
			    mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 1:
		P.planet = "AZAN"
		P.alien = "BIALAR"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(15)
		AS.defenseskill = rand.Intn(15)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {
			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}

	case 2:
		P.planet = "DASH"
		P.alien = "DENGAR"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 3:
		P.planet = "GARRUS"
		P.alien = "GILINA"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 4:
		P.planet = "JUBAL"
		P.alien = "KAIDAN"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 5:
		P.planet = "ZUBER"
		P.alien = "ELROW"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 6:
		P.planet = "BALTROS"
		P.alien = "KATERNA"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 7:
		P.planet = "LUDISC"
		P.alien = "MANTRA"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 8:
		P.planet = "LUDAN"
		P.alien = "KEZMO"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	case 9:
		P.planet = "JUBAL"
		P.alien = "KAIDAN"
		AS.alien = P.alien
		AS.attackskill = rand.Intn(14)
		AS.defenseskill = rand.Intn(14)
		if f==1 {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				mining(P,m)
				return
			}
		} else {

			if c < 5 {
				battleonplanet(TS, P, AS,f)
				return
			} else {
				trade(P,m)
				return
			}
		}
	}
}
func mining(P Planet, m int) {
	if m == 0 {
		fmt.Printf("\n\tYOU HAVE LANDED ON PLANET '%s'", P.planet)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tYou have encountered the alienspecies of the planet: %s", P.alien)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tThey seems to be peaceful towards your team.")
		time.Sleep(2 * time.Second)
	    fmt.Printf("\n\n\tLETS GET READY FOR MINING\n")
		time.Sleep(2 * time.Second)
	}
	if m==1 {
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\n\tLETS GET READY FOR MINING\n")
		time.Sleep(2 * time.Second)
	}
	k:=rand.Intn(6)
	b:=T.miningskill-k
	time.Sleep(2 * time.Second)
	fmt.Printf("\n\tWE HAVE GAINED %d GEMS\n",b*10)
	T.Gems=T.Gems+b*10
	time.Sleep(2 * time.Second)
	fmt.Printf("\n\tNEW GEMS-->%d\n",T.Gems)
	time.Sleep(2 * time.Second)
	return
}

func trade(P Planet, m int) {
	if m==0 {
		fmt.Printf("\n\tYOU HAVE LANDED ON PLANET '%s'", P.planet)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tYou have encountered the alienspecies of the planet: %s", P.alien)
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tThey seems to be peaceful towards your team.")
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tLETS START TRADING\n")
		time.Sleep(2 * time.Second)
	}
	if m==1 {
		time.Sleep(2 * time.Second)
		fmt.Printf("\n\tLETS START TRADING\n")
		time.Sleep(2 * time.Second)
	}
	fmt.Printf("\n\tORIGINAL FUEL-->%d\n",T.Fuel)
	fmt.Printf("\n\tORIGINAL GEMS-->%d\n",T.Gems)
	k:=rand.Intn(6)
	b:=T.tradingskill-k
	T.Fuel=T.Fuel+b*10
	time.Sleep(2 * time.Second)
	fmt.Printf("\n\tWE HAVE GAINED %d UNITS OF FUEL \n",b*10)
	time.Sleep(2 * time.Second)
	fmt.Printf("\n\tWE USED 10 GEMS\n ")
	T.Gems=T.Gems-10
	time.Sleep(2 * time.Second)
	fmt.Printf("\n\tNEW FUEL-->%d\n",T.Fuel)
	fmt.Printf("\n\tNEW GEMS-->%d\n",T.Gems)
	time.Sleep(2 * time.Second)
	return

}
func fuelrestore(){
	fmt.Printf("\n\tORIGINAL FUEL-->%d\n",T.Fuel)
	fmt.Printf("\n\tWE ARE GOING TO ADD 10 UNITS OF FUEL\n")
	T.Fuel=T.Fuel+10
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tWE HAVE ADDED EXTRA FUEL\n")
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tNEW FUEL-->%d\n",T.Fuel)
	time.Sleep(2*time.Second)
	return
}
func Gemsrestore() {
	fmt.Printf("\n\tORIGINAL Gems-->%d\n",T.Gems)
	fmt.Printf("\n\tWE ARE GOING TO ADD 10 UNITS OF Gems\n")
	T.Gems=T.Gems+10
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tWE HAVE ADDED EXTRA Gems\n")
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tNEW Gems-->%d\n",T.Gems)
	time.Sleep(2*time.Second)
	return
}
func battleonspaceship(TS TravelShip, ASP AlienSpaceship){
	var s, pa, pt int
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tYOU HAVE ENCOUNTERED THE ALIENSPACESHIP: %s\n",ASP.alien)
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tDO YOU WANT TO BATTLE IT OR GO BACK TO YOUR PLANET? \n")
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tREMEMBER, WINING THOSE BATTLE WILL GIVE YOU POINTS, AND THE MORE POINTS YOU HAVE" +
		",THE MORE EXITING PRIZES YOU CAN GET\n ")
	time.Sleep(2*time.Second)
	fmt.Printf("\n\tSO CHOOSE: 1.BATTLE" +
		                        "\n\t\t2.GO BACK")
	fmt.Printf("\n\tSELECT: ")
	fmt.Scan(&s)
	switch s {
	case 1:
		   time.Sleep(2*time.Second)
		   fmt.Printf("\n\t WE ARE READY FOR BATTLE\n")
		   time.Sleep(2*time.Second)
		   fmt.Printf("\n\tTHE ALIENSPACESHIP '%s' IS SHOWING SIGNS OF HOSTILITY\n",ASP.alien)
		   time.Sleep(2*time.Second)
		   fmt.Printf("\n\tITS GETTING READY FOR BATTLE\n")
		   time.Sleep(2*time.Second)
		   fmt.Printf("\n\tLET THE BATTLE BEGIN\n")
		   ASP.evasionskill=rand.Intn(14)
		   ASP.speed=rand.Intn(14)
		   ASP.attackskill=rand.Intn(14)
		   ASP.ArmourStrength=rand.Intn(14)
		   for i:=0;i<3;i++{
		   	fmt.Printf("\n\tROUND %d-->\n",i+1)
		   	time.Sleep(2 * time.Second)
		   	fmt.Printf("\n\t WE ARE ATTACKING\n")
		   	if ASP.evasionskill>TS.speed {
				time.Sleep(2*time.Second)
		   		fmt.Printf("\n\tTHE ALIENSPACESHIP DODGED YOUR ATTACK\n")
				time.Sleep(2*time.Second)
		   		pa+=10
		   		fmt.Printf("\n\tTHE ALIENSPACESHIP GAINED 10 VALOR POINTS\n")
				time.Sleep(2*time.Second)
		   		} else if ASP.evasionskill<TS.speed {
		   			if TS.attackskill>ASP.ArmourStrength {
						time.Sleep(2 * time.Second)
						fmt.Printf("\n\tTHE ALIENSPACESHIP TOOK DAMAGE FROM YOUR ATTACK\n")
						time.Sleep(2 * time.Second)
						pt += 10
						fmt.Printf("\n\tYOU GAINED 10 VALOR POINTS\n")
						time.Sleep(2*time.Second)
					} else if TS.attackskill<ASP.ArmourStrength{
						time.Sleep(2*time.Second)
						fmt.Printf("\n\tTHE ALIENSPACESHIP PREVENTED THE DAMAGE\n")
						time.Sleep(2*time.Second)
						pa+=10
						fmt.Printf("\n\tTHE ALIENSPACESHIP GAINED 10 VALOR POINTS\n")
						time.Sleep(2*time.Second)
				}
			    }
		   	fmt.Printf("\n\tTHE ALIENSPACESHIP IS ATTACKING\n")
		   	time.Sleep(2 * time.Second)
            if TS.evasionskill>ASP.speed {
				time.Sleep(2 * time.Second)
				fmt.Printf("\n\tWE DODGED THE ATTACK\n")
				time.Sleep(2 * time.Second)
				pt+=10
				fmt.Printf("\n\tYOU GAINED 10 VALOR POINTS\n")
				time.Sleep(2 * time.Second)
			} else if TS.evasionskill<ASP.speed {
				if ASP.attackskill>TS.ArmourStrength {
					time.Sleep(2 * time.Second)
					fmt.Printf("\n\tWE TOOK DAMAGE FROM THE ALIENSPACESHIP'S ATTACK\n")
					time.Sleep(2 * time.Second)
					pa+=10
					fmt.Printf("\n\tTHE ALIENSPACESHIP GAINED 10 VALOR POINTS\n")
					time.Sleep(2 * time.Second)
				} else if ASP.attackskill<TS.ArmourStrength {
					time.Sleep(2 * time.Second)
					fmt.Printf("\n\tWE PREVENTED THE DAMAGE\n")
					time.Sleep(2 * time.Second)
					pt+=10
					fmt.Printf("\n\tYOU GAINED 10 VALOR POINTS\n")
					time.Sleep(2 * time.Second)
				}
			}
            if pa>=pt{
				fmt.Printf("\n\tYOUR VALOR POINTS--->%d\n",pt)
				fmt.Printf("\n\tALIENSPACESHIP VALOR POINTS--->%d\n",pa)
            	time.Sleep(2 * time.Second)
				fmt.Printf("\n\tTHE ALIENSPACESHIP WON BY %d POINTS\n",pa-pt)
				time.Sleep(2 * time.Second)

			} else if pt>pa{
				fmt.Printf("\n\tYOUR VALOR POINTS--->%d\n",pt)
				fmt.Printf("\n\tALIENSPACESHIP VALOR POINTS--->%d\n",pa)
				time.Sleep(2 * time.Second)
				fmt.Printf("\n\tYOU WON BY %d POINTS\n",pt-pa)
				time.Sleep(2 * time.Second)
                T.VALOR+=10
			}

		   }

		   default:
		     time.Sleep(2*time.Second)
		     fmt.Printf("\n\tLETS GO BACK TO OUR PLANET\n")
		     return
		     }
}
