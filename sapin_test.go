package sapin

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleSapin_1() {
	fmt.Println(NewSapin(1).String())
	// Output:
	//    *
	//   ***
	//  *****
	// *******
	//    |
}

func ExampleSapin_3() {
	fmt.Println(NewSapin(3).String())
	// Output:
	//           *
	//          ***
	//         *****
	//        *******
	//         *****
	//        *******
	//       *********
	//      ***********
	//     *************
	//      ***********
	//     *************
	//    ***************
	//   *****************
	//  *******************
	// *********************
	//          |||
	//          |||
	//          |||
}

func ExampleSapin_presents() {
	sapin := NewSapin(4)
	sapin.AddPresents()
	fmt.Println(sapin.String())
	// Output:
	//               *
	//              ***
	//             *****
	//            *******
	//             *****
	//            *******
	//           *********
	//          ***********
	//         *************
	//          ***********
	//         *************
	//        ***************
	//       *****************
	//      *******************
	//     *********************
	//       *****************
	//      *******************
	//     *********************
	//    ***********************
	//   *************************
	//  ***************************
	// *****************************
	//             |||||
	//             |||||   _8_8_
	//             |||||  |  |  |_8_
	//             |||||  |__|__|___|
}

func ExampleSapin_fullballs() {
	sapin := NewSapin(3)
	sapin.AddBalls(100)
	fmt.Println(sapin)
	// Output:
	//           @
	//          @@@
	//         @@@@@
	//        @@@@@@@
	//         @@@@@
	//        @@@@@@@
	//       @@@@@@@@@
	//      @@@@@@@@@@@
	//     @@@@@@@@@@@@@
	//      @@@@@@@@@@@
	//     @@@@@@@@@@@@@
	//    @@@@@@@@@@@@@@@
	//   @@@@@@@@@@@@@@@@@
	//  @@@@@@@@@@@@@@@@@@@
	// @@@@@@@@@@@@@@@@@@@@@
	//          |||
	//          |||
	//          |||
}

func ExampleSapin_complex() {
	rand.Seed(42)
	sapin := NewSapin(3)
	sapin.AddBalls(4)
	sapin.AddStar()
	sapin.Emojize()
	fmt.Println(sapin)
	// Output:
	//           💛
	//          🎄🎄🎄
	//         🎄🎄🎄🎄🎄
	//        🎄🎄🎄🎄🎄🎄🎄
	//         🎄🔴🎄🎄🎄
	//        🎄🎄🎄🎄🎄🔴🎄
	//       🎄🎄🎄🔴🎄🎄🎄🎄🎄
	//      🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄
	//     🎄🎄🎄🎄🎄🎄🎄🎄🎄🔴🎄🎄🎄
	//      🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄
	//     🎄🎄🎄🎄🔴🎄🔴🎄🎄🎄🎄🎄🎄
	//    🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🔴
	//   🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄
	//  🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄
	// 🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄🎄
	//          🚪🚪🚪
	//          🚪🚪🚪
	//          🚪🚪🚪
}

func ExampleSapin_4balls() {
	rand.Seed(42)
	sapin := NewSapin(3)
	sapin.AddBalls(4)
	fmt.Println(sapin)
	// Output:
	//           *
	//          ***
	//         *****
	//        *******
	//         *@***
	//        *****@*
	//       ***@*****
	//      ***********
	//     *********@***
	//      ***********
	//     ****@*@******
	//    **************@
	//   *****************
	//  *******************
	// *********************
	//          |||
	//          |||
	//          |||
}

func ExampleSapin_garlands() {
	rand.Seed(41)
	sapin := NewSapin(3)
	sapin.AddGarlands(4)
	fmt.Println(sapin)
	// Output:
	//           *
	//          **~
	//         *~~**
	//        ~~****~
	//         ***~~
	//        **~~***
	//       *~~****~~
	//      ~~****~~***
	//     ***~~~~******
	//      **~~~~*****
	//     *~~****~~****
	//    ~~********~~***
	//   *************~~**
	//  ****************~~*
	// *******************~~
	//          |||
	//          |||
	//          |||
}

func ExampleSapin_10() {
	fmt.Println(NewSapin(10).String())
	// Output:
	//                                                   *
	//                                                  ***
	//                                                 *****
	//                                                *******
	//                                                 *****
	//                                                *******
	//                                               *********
	//                                              ***********
	//                                             *************
	//                                              ***********
	//                                             *************
	//                                            ***************
	//                                           *****************
	//                                          *******************
	//                                         *********************
	//                                           *****************
	//                                          *******************
	//                                         *********************
	//                                        ***********************
	//                                       *************************
	//                                      ***************************
	//                                     *****************************
	//                                       *************************
	//                                      ***************************
	//                                     *****************************
	//                                    *******************************
	//                                   *********************************
	//                                  ***********************************
	//                                 *************************************
	//                                ***************************************
	//                                   *********************************
	//                                  ***********************************
	//                                 *************************************
	//                                ***************************************
	//                               *****************************************
	//                              *******************************************
	//                             *********************************************
	//                            ***********************************************
	//                           *************************************************
	//                              *******************************************
	//                             *********************************************
	//                            ***********************************************
	//                           *************************************************
	//                          ***************************************************
	//                         *****************************************************
	//                        *******************************************************
	//                       *********************************************************
	//                      ***********************************************************
	//                     *************************************************************
	//                         *****************************************************
	//                        *******************************************************
	//                       *********************************************************
	//                      ***********************************************************
	//                     *************************************************************
	//                    ***************************************************************
	//                   *****************************************************************
	//                  *******************************************************************
	//                 *********************************************************************
	//                ***********************************************************************
	//               *************************************************************************
	//                   *****************************************************************
	//                  *******************************************************************
	//                 *********************************************************************
	//                ***********************************************************************
	//               *************************************************************************
	//              ***************************************************************************
	//             *****************************************************************************
	//            *******************************************************************************
	//           *********************************************************************************
	//          ***********************************************************************************
	//         *************************************************************************************
	//        ***************************************************************************************
	//             *****************************************************************************
	//            *******************************************************************************
	//           *********************************************************************************
	//          ***********************************************************************************
	//         *************************************************************************************
	//        ***************************************************************************************
	//       *****************************************************************************************
	//      *******************************************************************************************
	//     *********************************************************************************************
	//    ***********************************************************************************************
	//   *************************************************************************************************
	//  ***************************************************************************************************
	// *****************************************************************************************************
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
	//                                              |||||||||||
}

func TestSapin_String(t *testing.T) {
	Convey("Testing Sapin.String()", t, func() {
		Convey("size=1", func() {
			sapin := NewSapin(1)
			So(fmt.Sprintf("\n%s", sapin.String()), ShouldEqual, `
   *
  ***
 *****
*******
   |
`)
		})
		Convey("size=0", func() {
			sapin := NewSapin(0)
			So(sapin.String(), ShouldBeEmpty)
		})
		Convey("size=-1", func() {
			sapin := NewSapin(-1)
			So(sapin.String(), ShouldBeEmpty)
		})
		Convey("size=1 garlands=1", func() {
			sapinA := NewSapin(1)
			sapinA.AddGarlands(1)
			sapinB := NewSapin(1)
			So(sapinA.String(), ShouldEqual, sapinB.String())
		})
		Convey("size=5 garlands=10 with garlands outside of the sapin", func() {
			sapin := NewSapin(5)
			rand.Seed(42)
			sapin.AddGarlands(40)
			expected := `                   *
                  ***
                 ****~
                ***~~*~
                 ~~*~~
                **~~***
               ~~~***~~~
              ~~*~~~~~~~~
             ~~**~~~~~~***
              ~~~~~~~~~**
             ~~~~~~****~~~
            ~~~~~*~~~~~~*~~
           ~~~~~*~~~~~~~~**~
          ~~~~~*~~~~~~~~~~~**
         ~~****~~~~~~~~~~~~~~*
           ~~~~~*~~*~~~~~~~~
          *~~~~~~**~~*~~~~~*~
         ~~~*~~~~~~**~~*~~*~~*
        ~~*~~**~~~~~~**~~*~~*~~
       ~~~~~*~~**~~~~****~~*~~*~
      ~~~~~~~~*~~****~~****~~*~~*
     ~~*~~~~~~~~*~~****~~****~~*~~
       ~~*~~~~~*~~*~~****~~****~
      ***~~*~~*~~*~~*~~****~~****
     ~~~~**~~*~~*~~*~~*~~****~~***
    *~~~~****~~*~~*~~*~~*~~****~~**
   ~~****~~****~~*~~*~~*~~*~~****~~*
  ***~~****~~****~~*~~*~~*~~*~~****~~
 ******~~****~~****~~*~~*~~*~~*~~****~
*********~~****~~****~~*~~*~~*~~*~~****
                 |||||
                 |||||
                 |||||
                 |||||
                 |||||
`
			So(sapin.String(), ShouldEqual, expected)
		})
		Convey("size=3 garlands=2 balls=4 star colorize", func() {
			rand.Seed(42)
			sapin := NewSapin(3)
			sapin.AddStar()
			sapin.AddBalls(4)
			sapin.AddGarlands(2)
			sapin.ColorOpts = "bh"
			sapin.Colorize()
			expectedB64 := strings.TrimSpace(`
ICAgICAgICAgIBtbMTs5M20jG1swbQogICAgICAgICAbWzE7OTJtKhtbMG0bWzE7
OTJtKhtbMG0bWzE7OTJtKhtbMG0KICAgICAgICAbWzE7OTJtKhtbMG0bWzE7OTJt
KhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0KICAgICAg
IBtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1sw
bRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbQogICAgICAgIBtb
MTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5MW1AG1swbRtbMTs5Mm0qG1swbRtb
MTs5M21+G1swbQogICAgICAgG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzky
bSobWzBtG1sxOzkzbX4bWzBtG1sxOzkzbX4bWzBtG1sxOzkybSobWzBtG1sxOzkx
bUAbWzBtCiAgICAgIBtbMTs5M21+G1swbRtbMTs5M21+G1swbRtbMTs5M21+G1sw
bRtbMTs5M21+G1swbRtbMTs5MW1AG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1sw
bRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbQogICAgIBtbMTs5Mm0qG1swbRtbMTs5
M21+G1swbRtbMTs5M21+G1swbRtbMTs5M21+G1swbRtbMTs5M21+G1swbRtbMTs5
Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5Mm0qG1swbRtbMTs5
Mm0qG1swbRtbMTs5Mm0qG1swbQogICAgG1sxOzkzbX4bWzBtG1sxOzkzbX4bWzBt
G1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBt
G1sxOzkzbX4bWzBtG1sxOzkzbX4bWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBt
G1sxOzkxbUAbWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtCiAgICAgG1sxOzky
bSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzky
bSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkzbX4bWzBtG1sxOzkz
bX4bWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtCiAgICAbWzE7OTJtKhtbMG0b
WzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0b
WzE7OTFtQBtbMG0bWzE7OTJtKhtbMG0bWzE7OTFtQBtbMG0bWzE7OTJtKhtbMG0b
WzE7OTJtKhtbMG0bWzE7OTNtfhtbMG0bWzE7OTNtfhtbMG0bWzE7OTJtKhtbMG0K
ICAgG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSob
WzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSob
WzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSob
WzBtG1sxOzkybSobWzBtG1sxOzkzbX4bWzBtG1sxOzkzbX4bWzBtCiAgG1sxOzkx
bUAbWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzky
bSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzky
bSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzky
bSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkz
bX4bWzBtCiAbWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7
OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7
OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7
OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7
OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0bWzE7OTJtKhtbMG0KG1sx
OzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sx
OzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sx
OzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sx
OzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sx
OzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sxOzkybSobWzBtG1sx
OzkybSobWzBtCiAgICAgICAgIBtbMTszODs1OzkwbXwbWzBtG1sxOzM4OzU7OTBt
fBtbMG0bWzE7Mzg7NTs5MG18G1swbQogICAgICAgICAbWzE7Mzg7NTs5MG18G1sw
bRtbMTszODs1OzkwbXwbWzBtG1sxOzM4OzU7OTBtfBtbMG0KICAgICAgICAgG1sx
OzM4OzU7OTBtfBtbMG0bWzE7Mzg7NTs5MG18G1swbRtbMTszODs1OzkwbXwbWzBt
Cg==`)
			expected, err := base64.StdEncoding.DecodeString(expectedB64)
			So(err, ShouldBeNil)
			So([]byte(sapin.String()), ShouldResemble, expected)
		})
	})
}

func TestTriangularNumber(t *testing.T) {
	Convey("Testint triangularNumber()", t, func() {
		So(triangularNumber(3), ShouldEqual, 6)
		So(triangularNumber(4), ShouldEqual, 10)
		So(triangularNumber(1), ShouldEqual, 1)
		So(triangularNumber(0), ShouldEqual, 0)
	})
}
