package main

import (
	"fmt"
	"time"
)

func wait(t time.Duration) {
	time.Sleep(t * time.Second)
}

func main() {
	p := Printer{}
	p.createRequest(1, 2, "Up")
	p.createState(1, 3500, "MAINTENANCE", 1)
	p.upArrow("3")
	p.downArrow("1")
	p.doorOpen("1")
	p.doorClose("2")
	p.createArrival(8000)
	p.createPointing(5, []int{4, 5, 6, 78, 234, 3243}, []int{2, 3, 4, 5, 63, 225})
}

// Printer ...
type Printer struct {
	id int
}

func (p *Printer) positive(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func (p *Printer) createState(_id int, _floor int, _status string, _stop int) {
	wait(2)
	p.topBottomLine(_status)
	p.innerArrowElevatorLine(_status)
	p.emptyDoubleLine(_status)
	p.idLine(_id, _status)
	p.floorLine(_floor, _status)
	p.statusLine(_status)
	p.stopLine(_stop, _status)
	p.emptyDoubleLine(_status)
	p.innerArrowLine(_status)
	p.topBottomLine(_status)
}

func (p *Printer) createArrival(_floor int) {
	size := p.countStr(_floor)

	wait(2)
	p.topBottomLine(size)
	p.floorArivedLine(_floor)
	p.topBottomLine(size)
}

func (p *Printer) createRequest(_floor int, _stop int, _direction string) {
	wait(2)
	p.topBottomLine("2")
	p.innerArrowAndRequestLine()
	p.emptyDoubleLine("2")
	p.atFloorLine(_floor)
	p.floorRequestLine(_stop)
	p.directionLine(_direction)
	p.emptyDoubleLine("2")
	p.innerArrowLine("2")
	p.topBottomLine("2")
}

func (p *Printer) createPointing(_columnID int, _id []int, _points []int) {
	fmt.Println("")
	wait(2)

	p.columnSelectedLine(_columnID)
	fmt.Println("")

	for i := 0; i < len(_points); i++ {
		fmt.Println("		ELEVATOR", _id[i], "- HAS", _points[i], "pts")
		fmt.Println("")
	}

	p.bestElevatorLine(_id[0], _points[0])
	fmt.Println("")
}

func (p *Printer) bestElevatorLine(_id int, _points int) {
	fmt.Println("		THE BEST ELEVATOR IS ELEVATOR", _id, " WITH", _points, "pts")
}

func (p *Printer) columnSelectedLine(_id int) {
	fmt.Println("		THE SELECTED COLUMN IS COLUMN", _id)
}

func (p *Printer) countStr(n int) string {
	count := "0"

	if n >= 0 && n < 10 {
		count = "1"
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = "2"
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = "3"
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = "4"
	}

	return count
}

func (p *Printer) countInt(n int) int {
	count := 0

	if n >= 0 && n < 10 {
		count = 1
	} else if n < 100 && n > 9 || n < 0 && n > -10 {
		count = 2
	} else if n < 1000 && n > 99 || n < -10 && n > -100 {
		count = 3
	} else if n < 10000 && n > 999 || n < -100 && n > -1000 {
		count = 4
	}

	return count
}

func (p *Printer) topBottomLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		+------------------------------+")
	} else if _size == "MOVING" {
		fmt.Println("		+--------------------------------+")
	} else if _size == "1" {
		fmt.Println("		+-----------------------------------+")
	} else if _size == "2" {
		fmt.Println("		+------------------------------------+")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		+-------------------------------------+")
	} else if _size == "4" {
		fmt.Println("		+--------------------------------------+")
	}
}

func (p *Printer) innerArrowLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| +--->                  <---+ |")
	} else if _size == "MOVING" {
		fmt.Println("		| +--->                    <---+ |")
	} else if _size == "1" {
		fmt.Println("		| +--->                       <---+ |")
	} else if _size == "2" {
		fmt.Println("		| +--->                        <---+ |")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		| +--->                         <---+ |")
	} else if _size == "4" {
		fmt.Println("		| +--->                          <---+ |")
	}
}

func (p *Printer) innerArrowElevatorLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| +--->     Elevator     <---+ |")
	} else if _size == "MOVING" {
		fmt.Println("		| +--->      Elevator      <---+ |")
	} else if _size == "1" {
		fmt.Println("		| +--->        Elevator       <---+ |")
	} else if _size == "2" {
		fmt.Println("		| +--->        Elevator        <---+ |")
	} else if _size == "3" || _size == "MAINTENANCE" {
		fmt.Println("		| +--->         Elevator        <---+ |")
	} else if _size == "4" {
		fmt.Println("		| +--->         Elevator         <---+ |")
	}
}

func (p *Printer) elevatorLine(_id int) {
	count := p.countStr(_id)

	if count == "1" {
		fmt.Println("		  +--->      ELEVATOR", _id, "      <---+  ")
	} else if count == "2" {
		fmt.Println("		  +--->      ELEVATOR", _id, "     <---+  ")
	} else if count == "3" {
		fmt.Println("		  +--->     ELEVATOR", _id, "     <---+  ")
	} else if count == "4" {
		fmt.Println("		  +--->     ELEVATOR", _id, "    <---+  ")
	}
}

func (p *Printer) innerArrowAndRequestLine() {
	fmt.Println("		| +--->         REQUEST        <---+ |") // "2"
}

func (p *Printer) emptyDoubleLine(_size string) {
	if _size == "IDLE" {
		fmt.Println("		| |                          | |")
	} else if _size == "MOVING" {
		fmt.Println("		| |                            | |")
	} else if _size == "2" {
		fmt.Println("		| |                                | |")
	} else if _size == "MAINTENANCE" || _size == "3" {
		fmt.Println("		| |                                 | |")
	} else if _size == "4" {
		fmt.Println("		| |                                  | |")
	}
}

func (p *Printer) upArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| ▼      -/-\\-      ▼ |")
		fmt.Println("		|       /-/-\\-\\       |")
		fmt.Println("		|          |          |")
		fmt.Println("		| ▲       | |       ▲ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼          -          ▼ |")
		fmt.Println("		|          -/-\\-          |")
		fmt.Println("		|        -/-/-\\-\\-        |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼            -           ▼ |")
		fmt.Println("		|            -/-\\-           |")
		fmt.Println("		|          -/-/-\\-\\-         |")
		fmt.Println("		|         /-/-/-\\-\\-\\        |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|             | |            |")
		fmt.Println("		| ▲          | - |         ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

func (p *Printer) downArrow(_size string) {
	if _size == "1" {
		fmt.Println("		+---------------------+")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		| |                 | |")
		fmt.Println("		| ▼       | |       ▼ |")
		fmt.Println("		|          |          |")
		fmt.Println("		|       \\-\\-/-/       |")
		fmt.Println("		| ▲      -\\-/-      ▲ |")
		fmt.Println("		| |        -        | |")
		fmt.Println("		| +--->         <---+ |")
		fmt.Println("		+---------------------+")
	} else if _size == "2" {
		fmt.Println("		+-------------------------+")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| ▼                     ▼ |")
		fmt.Println("		|          | - |          |")
		fmt.Println("		|           | |           |")
		fmt.Println("		|        -\\-\\-/-/-        |")
		fmt.Println("		|          -\\-/-          |")
		fmt.Println("		| ▲          -          ▲ |")
		fmt.Println("		| |                     | |")
		fmt.Println("		| +--->             <---+ |")
		fmt.Println("		+-------------------------+")
	} else if _size == "3" {
		fmt.Println("		+----------------------------+")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| ▼          | - |         ▼ |")
		fmt.Println("		|             | |            |")
		fmt.Println("		|            |   |           |")
		fmt.Println("		|         \\-\\-\\-/-/-/        |")
		fmt.Println("		|          -\\-\\-/-/-         |")
		fmt.Println("		|            -\\-/-           |")
		fmt.Println("		| ▲            -           ▲ |")
		fmt.Println("		| |                        | |")
		fmt.Println("		| +--->                <---+ |")
		fmt.Println("		+----------------------------+")
	}
}

func (p *Printer) doorOpen(_size string) {
	if _size == "1" || _size == "2" {
		p.doorTopBottomLine(_size)
		p.doorMiddleLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.leftArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.openLine(_size)
		p.doorMiddleLine(_size)
		p.doorMiddleLine(_size)
		p.leftArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.doorTopBottomLine(_size)
	} else {
		fmt.Println("\n		The size entered for the door opening ain't good change it and try again !")
	}
}

func (p *Printer) doorClose(_size string) {
	if _size == "1" || _size == "2" {
		p.doorTopBottomLine(_size)
		p.doorMiddleLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.rightArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.closingLine(_size)
		p.doorMiddleLine(_size)
		p.doorMiddleLine(_size)
		p.rightArrowLine(_size)
		if _size == "2" {
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
			p.doorMiddleLine(_size)
		}
		p.doorMiddleLine(_size)
		p.doorTopBottomLine(_size)
	} else {
		fmt.Println("\n		The size entered for the door opening ain't good change it and try again !")
	}
}

func (p *Printer) leftArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  <<  <<  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   <<   <<   <<   ||      |")
	}
}

func (p *Printer) rightArrowLine(_size string) {
	if _size == "1" {
		fmt.Println("		|  >>  >>  ||   |")
	} else if _size == "2" {
		fmt.Println("		|   >>   >>   >>   ||      |")
	}
}

func (p *Printer) openLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Open   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Opening      ||      |")
	}
}

func (p *Printer) doorMiddleLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   	   ||   |")
	} else if _size == "2" {
		fmt.Println("		|     	           ||      |")
	}
}

func (p *Printer) doorTopBottomLine(_size string) {
	if _size == "1" {
		fmt.Println("		+----------++---+")
	} else if _size == "2" {
		fmt.Println("		+------------------++------+")
	}
}

func (p *Printer) closingLine(_size string) {
	if _size == "1" {
		fmt.Println("		|   Close  ||   |")
	} else if _size == "2" {
		fmt.Println("		|     Closing      ||      |")
	}
}

func (p *Printer) floorArivedLine(_floor int) {
	count := p.countStr(_floor)

	if count == "1" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "2" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "3" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	} else if count == "4" {
		fmt.Println("		| +--->  ARRIVE AT FLOOR :", _floor, " <---+ |")
	}
}

func (p *Printer) idLine(_id int, _status string) {
	count := p.countStr(_id)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "             | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "            | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "           | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "          | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "               | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "              | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "             | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "            | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       ID:", _id, "                    | |")
		} else if count == "2" {
			fmt.Println("		| |       ID:", _id, "                   | |")
		} else if count == "3" {
			fmt.Println("		| |       ID:", _id, "                  | |")
		} else if count == "4" {
			fmt.Println("		| |       ID:", _id, "                 | |")
		}
	}
}

func (p *Printer) directionLine(_direction string) {
	if _direction == "Up" {
		fmt.Println("		| |        DIRECTION:", "UP", "          | |")
	} else if _direction == "Down" {
		fmt.Println("		| |        DIRECTION:", "DOWN", "        | |")
	} else if _direction == "Stop" {
		fmt.Println("		| |        DIRECTION:", "STOP", "        | |")
	}
}

func (p *Printer) atFloorLine(_atFloor int) {
	count := p.countStr(_atFloor)

	if count == "1" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "            | |")
	} else if count == "2" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "           | |")
	} else if count == "3" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "          | |")
	} else if count == "4" {
		fmt.Println("		| |        AT FLOOR:", _atFloor, "         | |")
	}
}

func (p *Printer) floorRequestLine(_requestFloor int) {
	count := p.countStr(_requestFloor)

	if count == "1" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "     | |")
	} else if count == "2" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "    | |")
	} else if count == "3" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "   | |")
	} else if count == "4" {
		fmt.Println("		| |        FLOOR REQUESTED:", _requestFloor, "  | |")
	}
}

func (p *Printer) floorLine(_floor int, _status string) {
	count := p.countStr(_floor)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "        | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "       | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "            | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "           | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "          | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "         | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Floor:", _floor, "                 | |")
		} else if count == "2" {
			fmt.Println("		| |       Floor:", _floor, "                | |")
		} else if count == "3" {
			fmt.Println("		| |       Floor:", _floor, "               | |")
		} else if count == "4" {
			fmt.Println("		| |       Floor:", _floor, "              | |")
		}
	}
}

func (p *Printer) stopLine(_stop int, _status string) {
	count := p.countStr(_stop)

	if _status == "IDLE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "    | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "   | |")
		}
	} else if _status == "MOVING" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "        | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "       | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "      | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "     | |")
		}
	} else if _status == "MAINTENANCE" {
		if count == "1" {
			fmt.Println("		| |       Next-Stop:", _stop, "             | |")
		} else if count == "2" {
			fmt.Println("		| |       Next-Stop:", _stop, "            | |")
		} else if count == "3" {
			fmt.Println("		| |       Next-Stop:", _stop, "           | |")
		} else if count == "4" {
			fmt.Println("		| |       Next-Stop:", _stop, "          | |")
		}
	}
}

func (p *Printer) statusLine(_status string) {
	if _status == "IDLE" {
		fmt.Println("		| |       Status:", "IDLE", "      | |")
	} else if _status == "MOVING" {
		fmt.Println("		| |       Status:", "MOVING", "      | |")
	} else if _status == "MAINTENANCE" {
		fmt.Println("		| |       Status:", "MAINTENANCE", "      | |")
	}
}
