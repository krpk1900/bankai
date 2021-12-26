package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	text1 := "MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMBY''''TWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM{          ?MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM@!            dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMB'^~~?'TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM9!             .jMMMMMMMMMMMMMMMMMMMMMMMMMMBY'7!`    ~7TMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMK?77TTYYYBYYYYTT77?!~`            ?7MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMY!            .JgMMMMMMMMMMMMMMMMMf7<??!``                 ?TMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMN+                                  _TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#=`            .7''=<??7WMMMMMMMMMMMMN-               ....      .TMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMN,.                                 dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMB^                          _7WMMMMMMMMMNgJ,          (MMMM#~       (WMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMNm..              .....          .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM=                               (MMMMMMMMMM#%        .dMMMM#>         .MMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMNNmggggggggNNNMMMM}          dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMb              ..(x`             .dMMMMMMMMMD       .gMMMMM@!         .dMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMl          dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN;         .(gMMMY`           .(gMMMMMMMMMM=      .uMMMMM#=          (MMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMR          MMMMMMMMMMMMMMMMMMMMMMN- ?TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN&....JggMMMMMD!          .(gMMMMMMMMMMMD`      dMMMMMM=`         .dMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN.        .MMMMMMMMMMMMMMMMMMMMMMMN-   _TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNMMMMMMMMMY`        ..+NMMMMMMMMMMMM9!    .jm. ?TMB=          .dMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN.        (MMMMMMMMMMMMMMMMMMMMMMMMR.     7MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#=`      ..+gMMMMMMMMMMMMMMD-...(gMMMN+              .(MMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN_        (MMMMMMMMMMMMMMMMMMMMMMMMN_      .TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM9!   ..(igNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN-           .(MMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN`        jMMMMMMMMMMMMMMMMMMMMMMMMN!       .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM8u+ggNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,        .(MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN`        dMMMMMMMMMMMMMMMMMMMMMMMMN!       .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM9(MMMMMMMMMMMMMB'?!`?7TMMMMMMMMMMMMMMMMMMMMMMMN&.     (MMMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN`        dMMMMMMMMMMMMMMMMMMMMMMMMN!       (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMB! (MMMMMMMMMB=!          ?MMMMMMMMn_?TMMMMMMMMMM@!       ?7HMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#`        dMMMMMMMMMMMMMMMMMMMMMMMMN_       jMMMMMMMMMMMMMMMMMMMMMMMMMMMMMF`  .MMMMMM8!`   .ggNNm2    jMMMMMMMNo   7MMMMMMM@_           MMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#`        jMMMMMMMMMMMMMMMMMMMMMMMM#`       dMMMMMMMMMMMMMMMMMMMMMMMMMMMM@`   -MMMMY`     (MMMMMMb    .MMMMMMMM#    (WMMMMMMMMNP       (MMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMK         (MMMMMMMMMMMMMMMMMMMMMMMM$       .dMMMMMMMMMMMMMMMMMMMMMMMMMMM#!   .dMM= .j2    (MMMMMM#     dMMMMMMMD     .HMMMMMMMM$      .MMMBYTMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMHY=         -777TTYYYTYYYYYYYYHHHMMMH!       .MMMMMMMMMMMMMMMMMMMMMMMMMMM#^    .MY-.gMM=     ??7TMM#     JMMMMMM:       (MMMMMMMNl       `      (WMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMM: ?TMHYY''''7<?!``                                                       (MMMMMMMMMMMMMMMMMMMMMMMMMMM$     <!(7'!`            J#.    JMMMMM%         dMMMHY^^`               .MMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMK_                                                                        JMMMMMMMMMMMMMMMMMMMMMMMMMMB`                       .M#_    (MMMMM)                              ..JdMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMK_                                                                       .MMMMMMMMMMMMMMMMMMMMMMMMMMM:        ......        .(MB!     (MMMMM[                           .iNMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMK_             ...........(JJ++,         (qNNNNNNNNNNNNNmmggggga,        (MMMMMMMMMMMMMMMMMMMMMMMMMM%      + (MMMMMM{      .VY!       (MMMMMNe.          ..(dC       .&MMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMK_        (NNMMMMMMMMMMMMMMMMMMK         (MMMMMMMMMMMMMMMMMMMMM#{       .MMMMMMMMMMMMMMMMMMMMMMMMMM#:     (Nm._TMMMB!     _!    ..    (MMMMMMMMNgJJJJ+gNMMM8i.       dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMK_        dMMMMMMMMMMMMMMMMMMMM#.        (MMMMMMMMMMMMMMMMMMMMM#;      .dMMMMMMMMMMMMMMMMMMMMMMMMMM%      dMMN,              .(gN}    .MMMMMMMMMMMMMMMMMMMMMB>       ?7?!!~````~??TTMMMMMMMM\n" +
		"MMMMMMMMMMMM#~        dMMMMMMMMMMMMMMMMMMMM#`        .MMMMMMMMMMMMMMMMMMMMMMm.    .dMMMMMMMMMMMMMMMMMMMMMMMMMM#`     .dMMMMm.          .+MMMN}    .dMMMMMMMMMMMH9=` !`                           TMMMMMM\n" +
		"MMMMMMMMMMMMN~        dMMMMMMMMMMMMMMMMMMMMN.        .MMMMMMMMMMMMMMMMMMMMMMMNa-(&MMMMMMMMMMMMMMMMMMMMMMMMMMMN$      (MMMBY!        .(NMMMMM#}    .dB^7???!!`                                     MMMMMM\n" +
		"MMMMMMMMMMMMN!        dMMMMMMMMMMMMMMMMMMMMN`        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#`       ~`         .(gMMMMMMMM#:     dl               ....J-       (-......((JJ++ggNMMMMMM\n" +
		"MMMMMMMMMMMMM:        JMMMMMMMMMMMMMMMMMMMMN`        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMD                .JdMMMMMMMMMMM#!     dN-        ..(gNNMMMMN>       dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMM{        (MMMMMMMMMMMMMMMMMMMMN`        (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN:             .+NMMMMMMMMMMMMMM#`     dMN.  ..JgMMMMMMMMMMMN:      .dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMM}         MMMMMMMMMMMMMMMMMMMM#`        (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM`        ..JgMMMMMMMMMMMMMMMMMM@      dMMNNNMMMMMMMMMMMMMMMN!      .dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMM$         dMMMMMMMMMMMMMMMMMMM#`        (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMD       .dNMMMMMMMMMM#TMMMMMMMMMD      dMMMMMMMMMMMMMMMMMMMMN`      .MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMN.        dMMMMMMMMMMMMMMMMMMMK         (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM}       JMMMMMMMMMMMM#  ?TMMMMM#>      dMMMMMMMMMMMMMMMMMMMM#`      .MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMNx        jMMMMMMMMMMMMMMMMMMM$         (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM{       dMMMMMMMMMMMM#     ?TMM@       dMMMMMMMMMMMMMMMMMMMMK       (MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMN,       (MMMMMMMMMMMMMMMMMMM}         jMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMm,     .MMMMMMMMMMMMMMR.               dMMMMMMMMMMMMMMMMMMMM$       (MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMNagggNMMMMMMMMMMMMMMMMMMMM#:         dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNm.  .NMMMMMMMMMMMMMMMb.             .dMMMMMMMMMMMMMMMMMMMM}       .MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM%          dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNmgMMMMMMMMMMMMMMMMMMm.            (MMMMMMMMMMMMMMMMMMMMM:       .MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN{                                     (TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,          .dMMMMMMMMMMMMMMMMMMMMN!       .dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM[                                        _TMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNm,       .NMMMMMMMMMMMMMMMMMMMMMD`        dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN&                                         .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNJ...(gMMMMMMMMMMMMMMMMMMMMMMM$         dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMm,.         ......................      ..dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMr        .dMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNNNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN-       .MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN+      (MMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNm-...WMMMMMMMMMMMMMMMMMMMMMMM\n" +
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM\n"

	fmt.Printf("%s", text1)
	time.Sleep(1500 * time.Millisecond)

	// ここから分岐
	rand.Seed(time.Now().UnixNano())
	story_num := rand.Intn(3)
	switch story_num {
	case 0:
		story0 := "ltOOwV_dOltttttltllltlllllldf`  J0tttttlltltuV= .HSttlttttrtttrttrrttttttttttttttttttrtrtttttttttttrttttrttttttrttttttttttttttttttttlltltltltlltlllllllllllltltllllttlllllllllllllllllllllllltOrrrvzrrOt\n" +
			"AkX=` (0ltttllllltlllllllldY   .0lttllllludY!  .H0ttttltwQQHOtttttttttllltltlltlttttttttttttttttttrtttrttttttttttttttttttttttltllllltllllllllllllllllllllllltllllllllltlllllll===lllllllllllltrrt<jrrrrt\n" +
			"Y!   .Slllllllllllllll=llz=   .ZllllllluX=    .HHOllludY=(HSlttttttltlllltllllllllttltttttttttttttttttttttttlltltltttttlttlttlltltlllltlllllllllllllllllllllllllllllllllll=l==l=======llllllllttttzjwwwQ\n" +
			"     (Il=lllllv1ll=====1w!  ` J1===llzX=     .W9IlldK^  .H6llllltlllllllllllllltzv7!?7??!!~```````~????7?777<11zOOOOllllllllllllllllllllllzuOv1CQAz=======usOv<1yl============1?====l=====ll=lzwZ77!`.d9\n" +
			"     Zl======< +======1Z!    .z=====df!     .W6===df`  .XIlllllllllll===ll===llZU9w&(.....-dZVUVVVVVVVOOOOOOllllz+((..((<??71zzzlll=====X'` ..J9I=========1zVUI=========zG+(61?: 1=====?===uQV=   ..VIzz\n" +
			"     =<<<?<<!_-?<111=zZ!     (zzzzzzY`     .d6===z:   .dI===========l==============ll==lOVVIl=====ll=llllludXWzl==l==l==ludY= ..(dHR====zUTTOC?=====????????===?=?=????===?????>  <?????==d#^  .-VUAC!.J\n" +
			"          ........(-(-.    `.-(((((.       JVUVVU'   .dkAAAAQQaeszz=====================================zd=   4z==1zzz7=~..&WBU9Tzz1???????????????????????????????????????>???1.  (??????ZHa(d6?=TUVI1u\n" +
			"    (ZTC?<`  .v?????zW; `  .K???>??l    ` dD???1}    J1??===?==?=??==========================?========?1Z`   .TWWHaJJ+vTCz1??????????????>????>>???>??>?????>>>?>????????>>?>?>?z.  (>>>>???11?????1uZ=-\n" +
			"    Xz??<`  .+z?>>>>>j)    W$>>?>?+I     jK<???Z    JC????????????????=??==?=?===?=??===?=??==????=????v    .VU9Tz???????????>?>?>>>?>>?>>?>>>?>>>?>>>?>>>>>>>>>>>>>>>>>>>>>>>>>?<.  (>>>>>>???>>1dH+WB9\n" +
			"  `.WI>>!  .<=>?>>>>+d{   .9>>>>>>jD    .g3>?>zI..JVC?>>??>???????>????=??????????=??????????????????=dl   .0???>?>>>>>?>>???>>>?>>?>>?>>>>>>>>?>>>>>>>>>>>>>>>>>>>>>?>>>>>;>>>>>?<.  <>>>>>>>>?+TYC????\n" +
			"  .W3>>/ ` >?z>>>>>>+I   .k>>>>>>>jb .JYTC>>>>?+1?>>>>>>>??>>?>>???>???????=?????????????????????????1X~..-dI>?>??>>>??>>>?>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>_>>;;>;>;;>+?<   <>>>>>>>>>>>>>?>>\n" +
			"  (C>>>~  (?=z>>>>>>j:   dC>>>>>>>>?1<>>>>>>>>>>>>>>>>>>>>>>?>>>>>>???????????????????????????????>??dH0C1>>>?>>>>>>>>>>>>>>>>>?>>>>>>>>>>>>>>>>>>>>>>>;;>;>;>>>>;;>;+ .;;;;>>>;>+>;_  ->;>>>>>>>>>>>>>?\n" +
			". Z>>><   z=l<>>;;;>v   (D>;;>;>;;;>;>;>;>>>>>>>;>>>;>>>>>>>>>>>>>>>>>?>>????>>????>>>?????>>?>>???>>v1>>>>>>>>>>>>>>>>>>>>>>>>>>;;;;>>;>>>;>>;;>;>;;>;;;;>;;;>;;;>;;;.  ;;;;;;;>>z+:   ;;;;;;>;>;>>>>>>\n" +
			"Hk$>>;<  .Olv>;;;>;+:  .8;;>;;;;;;>;;;;;;;;;;;>>>;;>;;;>;;>>>>>>>;>>>>>>>>?>>?>>?>>?>>?>??>>?>>>>>>>>>>>>>>>>>;;>;;>>;;>>>>>>;;>;;;;;;;;;;;;;;;>;;;;;;;>;;;;;;;;;;;;;;_   ;;;;;;;<?l<   :;;;;;;;>;;>>>>>\n" +
			";<>;;;!  (kOI;;;;;;<  .9z;;;;;;>;>;;>;;;;;;>;;;;;>;;;;;;;>>;>>>;;>;>>>>>>>>>?>>>>>?>>>>>>>>>>>>>>>>>>>>>>>>;;;>;;;>;;>;;;;>;;>>;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;>;;;;;;;<-   :;:;;;<ll=.  (;;;;;;;>;;;;;>>\n" +
			";>;;;>`  dWwz>>;;;;< .$+<;;;;;;<<;;;;;>;;;;;;;;;;;;;;;;;>;;>;;;;;;;;;>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>;>;;;;>;;;;>;;>;;>;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:;;;<_:;;:;;<1_  (:;:;;<jwO;  .:::;;;;;;;;>;>;\n" +
			";;;;;<  .WWXz;;;;;<_.R+z;;;;;;`.;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;>;;;;;;;>>>>>>>>>>>>>>>>>>>>>;>;;;;;>;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:;;;:::::;:::::: _::;:::zl   ::::;<(WWl   ;:;;;;;;;;;;;;;\n" +
			";;;;;< `.@HK<;;;;;~_J11v;;;;<` (;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;>;;;;>>;;;;;>>>;>>>>;;>;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;:;;;;;:::::::::::;:;:::::;:::. .::;:<1?:  .:::::(HHR   ~:::;;;;;;;;;>;\n" +
			";;;;;~  (@@D;;:;;;_d:+lz::::`  <;::;:::;:;:::;;;;;;;:::;;;;;;;;;;;;;;;;;>;>>>>;;>;>;>;;>;;>;;;;;;;;;;;;;;;;;;;;;;;:;;;:;:;;::;;:;:;:::;::;:::;::::::::;::::::::<   ;:::j?<   :::::<H@H_  (::;:;;;;;;;;;;\n" +
			":;;;;`  J@@f;;:::<.P.zt<::;`  (<::::;:<;:::::::::;;:;:::;:;;:;;:;:;;;;;;;>;;;;>;>;;;;>;>;;;;;;;;;;;;;;;;;::;;::;::;:::;::::;:::;:::::;:;::::;:::::;::::;<::::::<_  _:::(+=.  ::::::WgH_` (:::::::::;:;;;\n" +
			";;;;;~  (@@D;;:;;;_d:+lz::::`  <;::;:::;:;:::;;;;;;;:::;;;;;;;;;;;;;;;;;>;>>>>;;>;>;>;;>;;>;;;;;;;;;;;;;;;;;;;;;;;:;;;:;:;;::;;:;:;:::;::;:::;::::::::;::::::::<   ;:::j?<   :::::<H@H_  (::;:;;;;;;;;;;\n" +
			":;;;;`  J@@f;;:::<.P.zt<::;`  (<::::;:<;:::::::::;;:;:::;:;;:;;:;:;;;;;;;>;;;;>;>;;;;>;>;;;;;;;;;;;;;;;;;::;;::;::;:::;::::;:::;:::::;:;::::;:::::;::::;<::::::<_  _:::(+=.  ::::::WgH_` (:::::::::;:;;;\n" +
			";::::   d@@$;::::_  (wO<::_  (+::;:::~(:::::;;:::::::;:;:::::;;:::::;:;;;;;;;;;;;;;;;;;;;;;;;;;;;:;;;:;::;::::;:;::;::::::::::::::::;::::::::::::::::::: ~:::::(;-  <::(zO-  _:::::d@g{`  :::::::::::;::\n" +
			"::::<   W@gI:::::~  zXk<::` .<?:::::~ (:::::::::::::::::::;:;;:::::::;;;;;;;;;;;;;;;;;;;;;;;;:;;:::::::;::::::::::::::::::::::::::::::::::::::::::::::::_ ~::::(<+. .::(OXr` (:::::d@@[`  :::::::::::;;:\n" +
			"::::<  .g@H<:::::   Xkk::_  (zv::::~  (:::::::::::::::::::::::::::::::;;;:;;;;:;;;;:;:;;;;:::;::;;::::::::::::::::::::::::::::::::::::::::::::::::::::::_  _::~(<+_  ::(dWR   ~~:::J@@L`  :::::::::::::;\n" +
			"::::!  .g@H;:::::  .HHD::~  1z>:::~  .<::::<:::::::::::::::::::::::::::::;::::;:::::;;::::;:::::::::::::::::::::::::::::::::::::::::::::::::~:~::~:<:~:::.  ~:::+z{  :::dHH_  ::~::Jg@b`  ~~::::::::::::\n" +
			"::~:~ `.g@#:::::_ `(gm$::~  zl>:~~  .(::::~~:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::~::::::::~::~::~::~_:~~(~.  ~::+tI  ~~:dgH;  ~~:~:(@@K`  _~::::::::::::\n" +
			"~::~` `J@@#<:~:~~  jgg$:~_ .ww<::` .>+:~~` :~:::~:::~::::::::::::::::::::::::;:::::::::::::::::::::::::~:~:~:~::~:~~:::::::::~::~::~~::~~:~:::~::~:~ _::(__  .~~jwX. _:~J@@[ `_~:~~(@@H`  .:~:~~::::::::\n" +
			":~~:` `J@@b<:~~:_  d@HI~:  (Wu<~~ `(zv:~~ .:~~~<~~:~:~~~:~::::::~:~::::::::::::::::::::::::::::~~:~~:~::~:~:~:~::~:~~~:~:~:~::~:~::~:~~::~:~~:~~~:~~. _:(<:_  :~(WW- .~:J@Mr` _~:~~(M@H_   ~:~~:~~~~::~:\n" +
			"~:~~  .d@@b<::~~`  d@H>~~ `jHR<~_  jv>~~ .;~:~`~::~:~:::~::~:::~:~:~~:::::::::::::::::::::::~~:~::~:~~~:~:~:~:~~:::::~:~~:~:~~:~:~~::~::~::~:~:_ ~::<. ~:<<:  ~:(g@)  ~:J@@b`  ~~:~<H@@_`  ~:~:~:~::~~::\n" +
			"~~:~   dgHD:~~~~  .W@H<~_  d@b:~~  dO>~` (<~:` :~:<~:~~:~:~:~~:~:~::~~:~::~::~::~::~::~::~:~::~:~~:~::~~:~~~:~:~~~~:~:~::~::~~::~:~~::~~~:~:<:~(  ~~<_ _~1+>  ~:(Hgr  ~:(@@b-  ~~~~~H@H{   ~~~~:~:~:~:~:\n"

		fmt.Printf("%s", story0)
		time.Sleep(1200 * time.Millisecond)

		fmt.Printf("~~~_   Xg@P~~~~~` .@@H<~~  W@D~:_ .Xw<~ .+>:~ <~~~_:~~~~:~~~:~~:~~::~:~:~~:~:~~:~:~~:~:~:~~:~~:~~:~~:~:~~::~~::~:~~:~~:~:~~:~:~~::~~:~:~:~~:_~~(_ _~;(. :jzO. (~(W@P` ~~(@gH_  ~~~~:X@H}   ~~~~~~~~~:~:~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_   W@@$~~:~~ `.g@H<~~ .g@D~~  (Wk<~ .z<_ (>~~ _~:(:~~~~:~:~~~:~~~:~:~:~:~:~:~~::~~:~::~:~:~~:~~~~~~~~~~~~:~~~::~:~:~~~:~~~~~~:~:~~~:_~~~_ :_;_ ~((~ ~(Xk-  ~(d@K  :~(@@M_  ~~~~(dM@r ` ~:~~:~~:~~~:~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_  .W@@I~~~~_  (g@@<~_ .@@$~~  JgK~_ (O~` 1>_ .<~~_~~:~~:~~~~:~~~:~:~~:~~~~:~:~~:~:~:~~~~:~~~~:~:~~:~~:~~~~~~~~~~~~~:~:~~:~~:~~:~~:~~~_:~< (~<< -(;:`~(Hg[` ~~dg#. ~~~H@@_` ~~~~(d@@$`  ~~~~~~~~~~:~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_  .W@@C~~~~_  (@@@~~` ,@@3~_ `dgK~_ df~ .O<_ +<~` ~~~~~~~~~~~~~~~~~:~~:~:~~~~~:~:~~~~~:~~~~~~~~~~~~:~~~~:~~:~~~:~~:~~~~~:~~~~~~~~:~~~ _~(_ _?>  (zc _(H@]  ~~d@H_ _~~H@@<` ~~~~(dM@D`` ~~~~~~~~~~~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_  .H@@I~~~~_ `J@@b~~  (gg>~_  H@D~` XP~ (v<`.=~_ (~~_~~~~~:~~~~~~~~~~:~~~:~:~~~~~~:~~~~~~:~~:~~~:~~~~~~~~~~~~~~~:~~~~~~~~~~~:~~~~~_~~~ ~(;.-++. ~dK .(H@b  ~~d@H_ -~~W@H{  _~~~_J@@b`  ~~~~~~~~~~:~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_  .g@g>~~~~_ `J@@b~~ `JHH>~_ .gH$~ .W$_`jw< (=_ .>~` ~~~~~~~~~~~:~~~~~~~~~~~~:~~~~~:~:~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~:~~:._~:_ (<_ jw_ ~XH  (W@b  ~~d@H: .~~d@@}   ~~~_J@@b`  ~~~~~~~~~~~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_  (g@H>~~~~_ `d@@D~~  J@N>~` .H@$~ ,@3_ dS~ jI_ (l_ <~~_~~~~~~~~~~~~~~~:~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~_~~< ~:< _<_ (W[ ~dH~ _d@H_ _~J@@{` ~_dHM[   ~~~_J@@b   .~~.~~~~~~~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~~_ `(@@M>~~~~` `d@@P~~  J@N<~  ,@@C~ (H>_ WS_ wI_ z>`(>~` ~__~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~:~~ _~~_ (<`~z> (gb _dH: _d@M_` ~j@@{  ~(d@@]   ~.~~(@@@_  .~.~~~~~~~~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~.~~ `(@@N>~~~~` .X@@$~~  d@N<_  J@@>_ jH> .gR_ y$_.z> (>` _~_~_~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~__~_~~ _~< (> _zy (H@  JH} ~d@H_  ~jM@{  ~(d@M]   ~.~~j@@@_  ~.~..~..~~~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~.~~ `(@@M>~~.~   W@M$~~  d@H<_ `J@H>_ dH< .@P_.W$_.w: d> (<`.~__~.~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~.~~~~__.~(_ _<`(z__dK (HH. (g[ ~d@H>  ~(@@}  ~(JH@D`  .~.~J@@K_  ..~..~..~.~~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~~._ `(@@M<~~~~   W@@r~~  d@#<~  J@M:_ dH~ .@$`.Hr`(X~.w! +: <_____.~.~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~.~~~~~.~~.~~..~.._~___< (<`_+ (H: dH (Wg_ (@r ~J@H{` ~(@@{  ~_J@@D`  ~...(@gK_  .~..~.~.~..~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(".~._ `(@HN<~.~_   W@@I~.` d@H~~  d@M<~ WK~ (@$`,@{ (W`.k`.z!.: !__~~~..~~~~~~~~~~~~~~~~~~~~~~~~.~~~~.~.~..~.~~.~.~..~..~..~~~__( <_(< (w_(g) JH..Wg{ (HD ~J@@r  ~(g@}  ~_JH@D`  ~.~.jg@K_  ..~....~.~.~.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("~.~_ `(@@M<~~._  .W@@I.~  d@#__  dMH~_ HK~ J@t ,g> JK ,W .O_(_.-~.`___~..~..~.~.~.~..~~..~.~..~~.~.~~.~..~.~..~.~.~..~..~_~___<._(_(<..H{-M[ JH~ Wg} (gP  (H@r  .(@@}  ~_JMMD ` .~.~J@@K.  .....~..~.~..\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(".~._ `(@@M<.~.~  .H@HI~.  d@K__  d@N~` MK. dg{ (H: dK JK -y (`(.~_  _~_~~.~~.~.~~.~~..~~.~~.~.~..~~.~.~~~.~.~..~..~~..~._~__.<(_<(< <_ H).g] (H: dg[ (Hb  (@@r  ~(@@{  ~_J@HP   ~..~jg@K`  ......~...~.~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("..~` `(@HN<..~.  .H@@l..` d@@__  d@H~ .gP_ W@> JH< dP db Jf z_<.__  .__~_~.~~~..~~.~~..~~..~.~.~~..~.....~..~..~.....~____. _~ _(.< ;_ W] H] (H{ dg] (HK` (@@$  ~(H@{  ~_d@@D`  ....JHgH`  ...........~.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...` `(M@M<....  .H@@l.~  dgK._  d@#_ .@P_ W@:`dH<`WF dF dP z <_._. __` ~_...~~.....~~..~..~......~..~..~..~.~.......___` ~___ _( <`>_ H] W] (m{ d@F`_HK  (@@$  .(@H{  ._dH@P.  ...~J@gK_  ........~....\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(".~.```(@@N:.~..` .H@@l..  dgK~_  d@#_ .@P_ HM< dH_ g] WF dF.I <_ __.._ `.~~....~~.~..~..~..~..~..~..~..~............_ ``` ___~ ~( +_>< W] WD (g) JgP _WH  (H@P  ~(@@l  ._d@@b`  ...~jg@H_  ...........~~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...` `(@@H>....  .H@gl..` dgK__  d@#_ .gF_.H@~`dH~ g] H] H].2.:( ~__._ `` -.~....~....~..~..~..~.................... ``  .____ ~(_+_(< dD WP (@) Jgb .WH  (@@$  _(M@r  ~_dHHb.  ....jg@#_  .............\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...` `(@@H:....  .W@gl..  dgK__  d@@_ ,@F_.Hg~ dH_ g] H] Ht({.!_ ~_.__ `   ..~.~....~..~............................  `  _~_~_ ~(_1_(<`dP Wb (g) JgK  WH  (@@D `_(@@r  ._dH@b.  ....jg@H_  .............\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...  `(@@N:....  .H@@l..  dgK__` d@@~`,HF_.Hg_ dH_ H].Mt M@(r.!:.~_~_   `  .........................................  ```.___-`~(_+:(>`db Xb`(g[ JHK  XH  (@@D  _(MHr  ._d@@b.  ....(g@H~  .```.`.......\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...  `(@@N>.`..  .W@gl..  dgK__  d@@. (@D`,HM_ d#_.H[.M] M)([(~<.__~~    ` ........................................`   ``.:_~..~(~+{(z`dK XK`(gr JgH  XH  (H@$  _(@Hr  ._dHMK.  ....jg@H~  .`..`.`......\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("..`  `(@@H>.``.` .Wggl`.  dgK__  d@@_ (@D`,HH_ d#_.H[.H[.M}j[(_<..___   `  ...........................`..```.```````   ``_<_~.._(~({(O d@`dK (gr J@H  XH` (@M$ `_(MMr  ._dHHK.  .`..jg@H_  .``.``.``....\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(".`.  `(@@H>`.`   .Wgg>..  dgK__ .d@@_`(@D`,HM_ WH~.H@.M[.M}dl(_:(.__       _.......................`.``.``..`..`.````  ` -~_~._~(~(r(k`dK dK`(@r JgH  XH` (@@$  _(@@r  ..dHMK.  `.`.j@MN_  .```.`.`.``..\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("`.`  `(g@H>.``.  .Wgg>`.` dgK__ .dH@_ (@D`(@M_ W#~.@[.Mr.M}d}+ <(_._    ` `.....................`.`.`..``.``.`.`..`` ` `  ~_~._~(_(r(K d@`dK (g$ JgH` XH` (@@$  _(@H$  ._d@@@.  .`..j@@N_   ``.```.`..`.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(".`.` `(g@H>``..  .W@g>`.  dgK__  d@H_`(@D`(@@_ W#~.H[.H[.M{d}+.:_--_       ...`.....`.........-(~_...`.`.`.``.````.` `  ` __~._~(_(P(H d@`dH (@P J@H. WH_ (H@r  _(@H$  __dHMK_  ``.`j@@N:   ````.`.`.`..\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("```` `(@@H>```.  .W@gl`.  JmK__  d@#_ (@D`-@H_ W#_.H[.@[.M:d{z :_  `      `.`.``.``.````.``.`.(Y3wv_``.``.`.``.`.`````    __.._~(:(PJH`dH dH (@P JgN_ XH_ (@@$  _(@M$  ..d@@K_  ```.j@HN:   `.```.```.``\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("```  `(g@H>```   .Wg@l`.` JgK__  d@H_ (@D.(@M_ H#_.H[,M].@:W:z <_. `   `   .`..`..`...`.`..`..d< <C.`.`.```.```.``````    -_.._~(>(PJN_dK dH_(gD (@N_ WN_ (@@$  _(HMD  . d@@K_  .``.jMHN:  -````````.``.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("``.` `(@gH>````  .W@g{`   JmK_.  dgH_ (@D`(@M..H#~.M[.M[.M;W:O <_  _`    ` .``.`..```..`..`..<>?!?_`.``.``.``.```.```   ` ___._~(>(PJN_d@ dH_(@D J@N_.WN_ (@@$  _(HMD  .-d@@K~  ```.(MMN>  _``````.``.`.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("```` `(@@H>```.  .W@gl`   JmK_. .d@H_ (@D (MN_.H#~.@),@[,M!W<w <_-.~ `     ````````````.`._` `__<<(__ `.``````.``````     ___._~(>(PJN`dK XH_(@D (@H~.WN_ (@@$  _(H@D  . d@@K_  ````(MHN>   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("```  `(@@H>``` ` .Wg@{`   JgK_.  d@H_`(@$ (gH..H#~.@),Mt,M!M<y <(-.~     ` .```````.`.``.(d/.   jcw  +```.```````.````  ` ___ _~->(PJN_d@ XH_(gD (@N~.WH_ (@@$  _(M@D  ` J@HH<  ````j@HN:   ````````.``.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("```  `(ggH{````  -WgH{` ` JgK_. .d@H_ J@$ (@H..HH_.M),H[,M!H:k.<:--~   ```. ``````.``. _!(@N~   JMH;.0 ``````.`````````   .__ _~({(PJN_XK WH_(gD (@H~.WH_ (@@$  _(MHD  . J@@#<   ```(MHN>   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `  `(@@H{ ``   (Wg@r`   JmK`. .d@#_`Jg$ (gH..HH_.g},M[.H_H~R.<:~-~  ```.` `````````-_`(@@N_   ,@@] d-```````````````` ` ___ _~-{(PJH X@ WH_(gD (@H~.WH~ (@@$  _(@HD    J@MH<  ````(@MN<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `(g@H{    ` (W@g}    JgK    d@H_`JH] (gH`.H#_.g},M@.N!H~K.::__~        ` `` ` ._``.HH@M)   (@Mb.(b_ ```` ````` `  `` `-  _~.{(bJH`XK WH_(gD`(gH_ WH~ (@@D  _(HHP    JM@H<      (@@N<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `(ggH}      (Wgmr    JmK    d@H. JgF (gH`.HH_,@},g[,N~H(K.:~                -~`   j@@H@] ` JHH@`.R -..                 ` `~.{(DJH`XK WH_(mD`,gH: WH~ (@M$   (M@P    J@HH< `    (M@M<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `(ggH}      (Wggr  ` JmK    d@H.`Jg] (@N .HH_,M},Mr,N_9``             `   .``    (HHH@@]  `?TT=  H-    ``.                  `(DJH_XK WH_(mR`,gN_ WH~ (@@$   ,@@P`   d@@H< `   `(MHM<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `jggH}    ` (Wgm}    JmK    d@H.`Jg] (HN .H#`,H}(m@ ?                .` `      .d@@HHH@@  `+~<(  Xr      (<                    _? WK WH_.mD`,gH! WH_ (@@P   -@Hb`   J@MH>  `   j@@M<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `  `jggH}      _Wgg}    JqK`  .dHN_ Jm] JgH .@#`,H}                      `      .(@@@@@@H@)   N(Hk_ Jb                                 `WH_.qR`,gH: WH_ (@@P   (@@b.   J@@H> `    j@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `j@@H}      (Wggl    JqK`  .dgH.`Jg] JH# ,H8                                 `?HMHHHH@M}  .MMHM_ ,f                                     .HS ,gH: WH: (@H$   -@@b`   dM@H> `    j@HM<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   ` .zggH}    ` (Wgg}    JqH.  .dgH_ Jg] JHH                         `.             ?HMH@@N~  ,@H@M{ .                                          .HH: WH: (@gR   (@Hb.   d@@H>  `  `j@@M<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   ` .dg@H}      (Wgg}    JqH.  .dgH_ Jm]                                              -?77=   (HHM@[                                                 WH{ (gg0   (@@K.   dg@H> `    j@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     .z@@H}    ` (Wgg}  `.dmH.  .dHH`  ``                                               _!     J@H9M]                                                 ?7! (gg0`  (@@b_   dg@H>`     (H@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     .zggH}      (Wgm}   .dmH_  .?''                                                          `dM@$HD  .                                                  (''=`  (@@R_   dH@H> ``   (@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   _  zg@H}  `   (Wgg}  `.dqH`                                                        `  `     WMH{gP  _                                                         (ggK_   dg@H>     `(@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   _  w@@H}    ` (Wggl `  `!!                             `                                   .H#XjHP                                                            .77^`  `dggH>` `   (@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   _  wg@H}    ` (WmH}                                                 `           `          .@$d?(D                                                                    dmgH> `    (@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   _  vg@H}      ?7''^                                                               -   `  ``,@NW+-F` . `                                                               ?TTT! `    (@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      umgH}                                                                       ` `_?_,. .  ,@HKj(F (!                                                                           `j@@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     -XggH}                                                                           _..!~(_ J@@Nddb.`                   `                                                         jH@H<   ````````````\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    ` 7YYY>                                               `            `              ,H@gHa-.d@@H@@gY`   `                 `                                                       ?'''!   ````````````\n")

	case 1:
		story1 := "                         `    `  ` `\n" +
			"   `   `  `  `  `  `  ` ``(.`  `.V^`  `               `    `   `     `   `     `   `     `   `\n" +
			"                          d~ ` ./              `    `             `         `         `\n" +
			"  `   `  `  `    `  `     ,_` `Z~           `           `                                                                                                                                              `\n" +
			"              `       `   .[`  .       `         `         `   `     `   `     `   `     `   `\n" +
			"       `           `    ` `1   (                    `             `         `         `                                                                         `   `   `\n" +
			"  `   `   `  `  `          .%` (.`   `    `            `       `         `         `          `                                                              `  `` ` ` `                               `\n" +
			"                    `  `   `(-  _              ``         `          `         `         `                                                                 ` ../?!```?^+,``\n" +
			"   `    `  `     `         ` (, < `         `   `.  `        `    `       `         `       `                                                          `   .C```    ``   1.`\n" +
			"      `      `       `        (o1.     `               `        `        `         `                                                                   ``.Y `     ``.(..`J#o.``                        `\n" +
			"  `             `  `      `    -;(+  `       `  .`        `          `      `  `      `  `   `                                                          .l      `  (MNHXo1MME `\n" +
			"       `  `  `        `      `  k db        `  .`   `        `    `                                                                                     ,:. `     .mHmHYdMWS?,   `\n" +
			"                              ` Z.$%  `      ` '       `       `        `        `  `         `                                                         dY 1.   `.g@HHkYdKWWb1                        `\n" +
			"  `  `   `  `   ` `  `    `    .0,>}`            `        `          `    `  `          `  `                                                           `742 1, .?!j(M#HHR4mW@=\n" +
			"        `              `    ` `.1o3}     `          `           `                     `                                                                 I(?! .Y.`.HHNW@KWMv7M1.  `\n" +
			"             `     `         `(:.t(~   `    `          `   `      `            `   `         `                                                          XnJXvTY(7YHJJg@DW@mW#(`                        `\n" +
			"  `  `   `      `     `    `  ._./2            `              `      `   `  `            `                                                          `` ,4J>d7?1dZI&BaHMmgd@HHs\n" +
			"       `    `       `         J_/?   `             `  `                            `  `                                                             `  e?=??!++J,X9gK4HM#6c<WM-\n" +
			"              `  `          ` (? c                       `     `  `      `     `             `                                                    ` ` J`   .....`_.M%      .nUl `                     `\n" +
			"  `   `  `           `  `  ` `i` '`         `  `           `         `      `            `                                                          .v_.`.Xl !~(.=!J#<1cz?G..bV_ `\n" +
			"           `      `           WG,`     `          `  `  `                 `      `  `         `                                                  ``.>(JL ZY=7?^!`  JMNUdHe>t CR}`\n" +
			"   `         `  `     `      `JyR  `                          `  `       `     `      `      `                                                 ` ` J--+64i?&  .J!`` (Hb1iJ(( _U) `  `                  `\n" +
			"      `  `         `         .#H$`        `    `                    `                    `                                                 `` ` .JH'.~_?J:/=(dV1m, QHXk.vTyT:.(W,      `\n" +
			"  `        `           `  `  ._Y(.     `            `   `  `    `      `    `      `                                                        ` .W@M@ .dGi(<(K^ `.MMMRdH ?/^a(t,Id@H,.`\n" +
			"       `     `  `   `        .!l1                `                             `      `     `                                     ` `   ``  .@@@@@% ,<3,`.@     (WY(HNX//l 1X !X@@@H,``  `            `\n" +
			"          `       `           2!/`          `         `       `   `  `   `          `    `    `                                `   `  `` .JH@@@@@@rSJWB(j#t>`   ,.%!??6@,,3d(:`X@@@@@N+   `  `\n" +
			"  `  `               `       `nJ!    `         `    `      `                `      `                                            `  ..(&W@@@@@@@@@MH@b+dJ,1d~ ` .j@d+  4gI_,rjmlJ@@@@@@@Mm,.` `  `\n" +
			"       `   `  `  `     `  `   fl                        `       `              `                                            ` ` `.W@@@@@@@@@@@@@@@@H<KSSN_-...(dD;weh..8G(.Xkt.W@@@@@@@@@@@N, ` ` `    `\n" +
			"         `          `       `.3l  `    `  `                         `   `                `   `                            `    .+M@@@@@@@@@@@@@@@@HD X,1] !~(HW$J`??j- +&<-(X-`w@@@@@@@@@@@M@@@m,  `\n" +
			"  `   `      `               .m)`              `  `  `     `   `          `       `   `                                     `.(@@@@@@@@@@@@@@@@@@@#  R~J{._,H$kd! ,6$I _?o~,(> J@@@@@@@@@@@@@@@@@Nmg, `\n" +
			"                `  `  `   ` `.cl            `           `         `         `  `             `                        `  ` .J@@@@@@@@@@@@@@@@@@@@HP`.HKd`` .VdNC `,=(n `(j<,.~ d@@@@@@@@@@@@@@@@@@@@@N..\n" +
			"   `    `  `               ` :d?`      `                             `   `         `     `                              `  J@@@@@@@@@@@@@@@@@@@@@HML,WH$,-+-.#1! JSwl,r_.JDhr!.@@@@@@@@@@@@@@@@@@@@@M@@h\n" +
			"      `     `   `    `     `, {<    `          `   `  `    `  `                       `       `                          `.M@@@@@@@@@@@@@@@@@@@@@@@N,HT~-(z-dX%``+(d??1_ +53I_-@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"  `               `     `  .+. .. `         `                   `         `   `              `                         `  J@@@@@@@@@@@@@@@@@@@@@@@@@HYj,($-.@J` `j?J0z(! v,+V-.@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"       `  `  `           ` .Zh .l`                      `         `  `           `  `    `                            ` `,@@@@@@@@@@@@@@@@@@@@@@@@@@@}_l++J5Z    ?Qc(Wvx`((<J<.M@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"                `  `  `     /b7i! `    `         `  `      `   `        `   `         `                            `    (@@@@@@@@@@@@@@@@@@@@@@@@@@@@{j@.gt11.  `JWnXy(W<4<dCv J@@@@@@@@@@@@@@@@@@@@@@@#\n" +
			"  `  `   `    `              ,x7.i. `       `         `                        `           `  `                     ` .J@@@@@@@@@@@@@@@@@@@@@@@@@@@@@>jz<R(>:   `(1(kvs+?8`1+l (H@@@@@@@@@@@@@@@@@@@@@@b\n" +
			"           `      `  `        j,_-I.l`         `                  `  `   `         `     `                          `.H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@MPj{.>d<~   `.?-d+Wr(_`(mk`` w@@@@@@@@@@@@@@@@@@@@@@\n" +
			"      `                `    `  4y<_(: `            `    `  `  `             `         `                         ``  .H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Nj(l,j]1:`   ((-ZJJ8.4,(Jd{`.J@@@@@@@@@@@@@@@@@@@@@@\n" +
			"  `     `  `  `  `         `  .dXI_(.    `  `         `         `              `    `        `                    `.@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@N,Zl2.].%. ``.v1(JG8,neV2!_q@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"                   `  `      `nrdI(J:``          `                  `   `                `                      ``.W@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@N b(:(.Jr`   .cJv<,36.J-.L.W@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"   `  `  `  `   `            .HlvT%-_               `     `    `          `       `           `              ```.dM@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@# Zd{_ (_   `.jIzJ( ,.ZM@@MH@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"                    `     ` `,@GJx(&z?         `        `         `         `  `      `      `            `   .(@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@MN{.}(H..,_    `?z9_> ,_WQH@@@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"  `       `  `   `    `      1iOS>`         `                `       `   `          `    `                 `.d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@bk+=}`j`` `     ,.f(` [??THMM@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"      `  `                `  Tghql`    `         ```.`     `    `                                          ` .@@@@@@@@@@@@@@@@@@@@@HMW@@H@@@@@@@@@XjU!) v  `       [>-..t  ,` ?@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"              `   `  `     `  (nV}`                 <   `                 `   `  `          `               J@@@@@@@@@@@@@@@@@@@@@@@fdMH@@@@@@@@@Hkv`,(.>`,`       <!},,{ J-i..M@H@@@@@@@@@@@@@@@@@@@@@@\n" +
			"  `    `   `           `     ` ?A+ `      `    `                  `  `              `  `      `            .@@@@@@@@@@@@@@@@@@@@@@@@F.H@@@@@@@@@@@}l..-W[ C``      .J_-(z+...$d@@@@@@@@@@@@@@@@@@@@@@@@@\n" +
			"      `         `              ./O:    `                      `         `   `            `             ```.HH@@@@@@@@@@@@@@@@@@@@@@@N M@@@@@@@@@@Mbkb-W4t`>      `  '_(.oGi<(YM@@@@@@@@@@@H@@H@M@H@@@@@@\n" +
			"   `      `  `    `  `          </u-  `     `         `    `    `              `   `         `            (@@@@@@@@@@@@@@@@@@@@@@@@@N d@@@@@@@@@@Mbl` ^!, (``      ` ({(_dHk0e(@@@@@@@@@@@@@@O(~?M@H@@@@\n" +
			"                               `.- z.`         `    `               `    `            `                 `(@@@@@@@@@@@@@@@@@@@@@@@@H@@pd@@@@@@@@@@@NJ. l ,(]          .:<^`((?d@@@@@@@@@@@@@@@:z,.M@@@@@@\n" +
			"  `   `  `    `  `    `   `     .G, }`                  `      `            `            `    `   ` ``   d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Ndk`-iJ < `    `  `./{`.J=(WW@@@@@@@@@H@@@@{ T>W@@@@@@\n" +
			"           `       `         `   %j(_`  `   `    `         `      `   `        `  `                    .d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@N.j>_ ?-[ `   ` , ` .` ?`,1/M@@@@@@@@@@@@H@Mx.= 4@@@@@\n" +
			"       `               `       `.! Jo               `                    `          `      `         ,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@Pi.(.-1<]     ` 2~ ,.  `.2i.d@@@@@@@@@@@H@@@@a.`,@@@@@\n" +
			"  `       `  `  `   `         ` +_ ??          `  ` `  `     `  `    `      `         `  `    `   ` .d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@H@@@H@@Y6J `Js.|>`     .vs,1<.jSu>(JM@@@@@@@@H@@rY4@MP!`.@@@@@\n" +
			"     `                     `  ` ?} -}    `  `     ` 1     `               `    `                  ``(@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@M@@@@@@@M($.(2,1.l ``    ,Dj:(c(4D?[-H@@@@@@@@@@@@D.d@#`` .@@@@@\n" +
			"       `   `    `  `  `        ._!,h} `             (             `                `               .M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@bM@@@@@@@N.i ..,.}  `  ` ~?I-(}  1.(d@@@@@@@@@@@@@} T^ `  ?W@@@@\n" +
			"  `      `    `              ` .r((X>                 `    `   `     `   `            `  `   ` ` `.W@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@H@@DM@@@@@@@@@N/,-_(}  {    !((.,} `.j!.M@@@@@@@@@H@N         /M@@@\n" +
			"      `           `    `      `.r 7_<`         `        `                   `  `    `           `.W@H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@MH?@@@@@@@@@@@h,.~_$` {   .-jdy<}  (jZ=MH@@@@@@@@@@N          HM@@\n" +
			"   `       `        `     `     ~ j~<``     `      `           `  `                         `  `.M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@B.@@@@@@@@@@@@>I-Zz  !    >.Jk_vc4vjS<M@@@@@@@@@@@N`         ?@@@\n" +
			"        `    `  `            ` `( .__                 `    `         `  `         `      `     `,@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@H@hd@@H@@@@@@@@Ml$,Wc, _`  (.7:- aJ;JWHldW@@@@@@@@H@N.  `      `?TM\n"

		fmt.Printf("%s", story1)
		time.Sleep(1200 * time.Millisecond)

		fmt.Printf("  `               `   `         ._  _    `     `                          `  `        `          M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@NH@@@@@@@@@@@@MtvC,I: _   ..>-<.@Nd@~l(PJbW@@@@@@@@@@o`   `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `  `  `                   .-.  `      `     `     `     `  `             `   `         `   M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@M#9HW@@@@@@@@@@H@MA&j2jk!.!   ~-_ :.?iw (  7dhg@@@H@@@@@@@N,`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `  `  `  `     `  .<` +`              `      `        `    `               `    `  d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@l.@@H@@@@@@@@@@@@@NJ!ZS  `   . ./(( ,?(; .QMH@@@@@@HHH@@M@M;\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `    `  `               `    `,<_ _                 `         `           `       `          ` .H@@@@@@@@@@@@@@@@@@@@@@@@@@@@H@@N`?WMM@@H@@@@@@@@@@K(<<6   `  !..C((> ` -JH@H@@@@@@M$svH@MHY!`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `      `    `           ` 3(_.c `     `  `        `            `         `   `             `,M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@N. Ui._7THH@@H@@@@@P-,>+`_, .  (>~ (   .(,W@M@@@@@@HJd_,HM{   `      `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `            `    `          +(`<     `         `       `  `   `     `             `  `  `   `.fT@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@L.io._?<-.?TTY^H@@M. !Z`~~`(?`-~. X.<z~/.j@@@@@@H@MkzuJ!. `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("        `  `          `      `` cC,~                  `                   `                   `  (@ W@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#(<1I-/7T9?=<zyxM@Ms-.C.~``._ _._.{ _?!,H@@@W@@@.=-~J7' (.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `         `    `  `     `    .>`(              `             `     `      `  `  `            `.f .H@@@@@@@@@@@@@@@@@@@@@@@@@@H@@Hlz.<..</?7<CI...z@@@@#_...    (( x--I -.H@@@ C!~>.umVu-.J `         `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `  `    `               `((.+.   `       `    `   `  `      `      `          `  `   `    d!.@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#(>dx..?1.+??-..`.-W@@M#/!~-  . ._:I&.`.@HH@Mb<j.r~?<_~(./.:.  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                      `       `(Z d_        `                                                ` .@ .QMH@@@@@@@@@@@@@@@@@@@@@@@@@@@@J@@@@@@Na.-<<???>(M@@@@$  `  . .__v?1. j@@@M1+?(...+_-!  .dhR\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `    `   `    `  `      `  ` (,6W)  `                       `      `      `  `         `  ` .F  T^!TH@@@H@@@H@@H@@@@@@@H@@@HH@@H@@@@@@@@@@@@MHHHHMH@@@MH/--  . <~1) ?;.d@@@8_(?^!~! `` ,`1dT!`       `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `      `        `       `-)?Ml`            `  `  `   `    `        `         `  `     `.f      `?M@@@@#= ` W@@@@@@@@@@#` ?M@@@@@@@@@@@@@@@@@#r4M@@MH].(....1.j_>,-d@@@@Hgga..,..d!( `.dn`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `      `       `            ._ j[      `    `                  `                      `  `,`   ` -`.4M@@M[    d@@@@@@@@@@b` J@@@@@@@@@@@@@@@@@@V(d?W@MHj?: I.((v+ vJH@@@M@@M@@@@@#HmHA-``~dm,\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `     `  `  ` `  .` (h `      `           `      `     `   `  `  `    `    ` .2     `? / .@@@Ml    dH@@@@@@H@#'.d@@@@@@@@@@@@@@@@@@Ht(j[ I41J@l < Y(<H@@@@@MT7@MJH@M@HNdeQh+.  dMJ.`     `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `   `  `       `           ` n ` ?;`              `      `                              .@.-    `./  `W@@#`   `.H@@@@@@@@_`(M@@@@@@@@@@@@@@@@@@bd(}>` j1(r(_!.-j-H@@@@HD `(MM@@@@@@HHNHb,  (dt\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("           `                  `(l ` x`  `        `                `               `   ` `.= Vf  ;,~v    4H#'`     T@MM@MMY^' (M@@@@@@@@@@@@@@@@H@NI+`?  ,  n-_..(_,MM@@@f`(: W@@@@@@@HNHMW[` (D    `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       `     `  `  `  `   `    .6.. T`      `         `      `  `    `  `  ` ` ``  ` ` `.:  (#(Cj`.^`  ` ?`      `  ?U@@@  ` vHYH@@@@@@@@@@@@@M@Mv(1?_``. `Z:..-.!J@@EM9~`(:.d@@@@@@HNdHHdF`.(c        `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `       `                  `   6.+J-.  `     `        `  `                 ` .1.`` .?=` `.,-zJ(d_`                ``T9` `..JJRC?MH@@@@@@@@@@@@Nr`(+<((-.`+:--<.H@@#C?U(<?@@dM@@@HMH@NKHWMJo<j\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `            `  `           (G,XK~             `                     `   `> |`.=` ` . .`..,v=?-      ```````` ....Jmd1XJ?HH,.jdHH@@@@@@@@@@@h, .Xm(y_ O:..I-@@#1!.J^<.5Yj@@@@@@@@@HuXHWf_,~`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       `   `  `        `        `` O#!                         `  `  `        `.;1.l ..x:`.1?.X3`,2{`    ...-uudzWW8SJ1db.Hu(Z6WZWQd@@@@@@@@@@@@@MM,(??9?` o!_ (HMM1^ .>.<(R^dH@@@@@@@HTWdb,!.[l       `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `      `      `   `        `    .J$`      `  `   `  `   `              `     `o C} .4[ ....(GJUXkedmudqK$J@`(Yk.JmHWWHvXXhZYugH@@@@@@@@@@@@@@@@@@N, .>..ld+((d@@@-.Z+4(>3.(MM@@@@@@@H@HNQJi.`,.  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `      `            `   ``.+TH;   `                       `     `   ` `..J?.(J&,`OHQWH#HHs!?4V==n/_.dNQsUH8QHhU=7?!  jHH@@@@@@@@@@@@@@@@@@@@@@#J/XdGd@K(W@@@b,(+/U+wg@@@@@@@@@@@@@@@@Nl   ;`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `              `  `        `.@ .W} ``          `  ` ` ` ```  `  `  `  ....(_`.,.d@Hm#ddBO+v?nvH/<JGWmgW9TV77=~    `    J@@@@@@@@@@@@@@@@@@@@@@@@@h/=vMxW@@@@@@@b>;.JH@@@MMMH@@@@@@@@@@@@Mh-  w      `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("        `  `    `      `   `  `[  .H'` ?>`  `     ` ....JJ?!?1, ...JJWdY7=`<-Q; `1 Jx(+S_d(pv?=,9JR=Bh77?! `              W@@@@@@@@@@@@@@@@@@@@@@@@@@N  ?d@@@@@@@@b-?_QM@M@H@@@@@@@@@@@@@@@@@_  (_`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `         `                  ,?t.6m..,8,.` `` `` .c..?!-dWf=<?~.._  (,1vC?C&L   l.dL  Jj .c...l.d{`                    .H@@@@@@@@@@@@@@@@@@@@@@@@@@N .d@@@@@@#WB5- dW#zMW@@@@@@@@@@@@@@@@@ML` ~i`  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `  `    `  ` ```  ` ` `` IQ. .U `(d` <1zwe,/JZZ1w>&z=<s/<4C.c:_1__~...(Y(-.`._(?``(dBkn(B'??!                     `.M@@@@@@@@@@@@@@@@@@@@@@@@@@Mo?X@@@@@@@Hl `.WKW#UdM@@@@@@@@@@@@@@@@@N. 1j.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("          `    ` .(1zzI>(>Y??-+OXMQW9VWjK?` ?OWBCjJiyzvuJC(J.x.J=-..-4Ak9??^??!???`i.?!``` ` `                         `.d@@@@@@@@@@@@@@@@@@@@@@@@@@@N~Ud@@@@@@HMo ,_WPMhd@H@@@@@@@@@@@@@@@@@M[` ;)\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(" ``` `` ` `....<!-..Jv<v=+x=??==!....1(J-...ZdH7vv&((mZ,(aasSz,?!`   `             .J~                                  .@@@@@@@@@@@@@@@@@@@@@@@@@@@H@,rO@@@@H@bJ>J JX(#W@@@@@@@@@@@@@@@@@@@@@N.`1(,` `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("...<?!_(<<(+(&&-(^``.......--C_~<Je-J-  ...v7wuJZ^7???~``    `                     ``    `                              .@@@@@@@@@@@@@@@@@@@@@@@@@@@@HH!(@@@@@@NJ2<v w@Nd@@@@@@@@@@@@@@@@@@@@@@F- ?b `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("T<>~`<(.....,I.-.uJx.-JY_<?(...J&wH@777~` ``  `                          `                                          `  `.@@@@@@@@@@@@@@@@@@@@@@@@@@@Mtj[.@@@@@@@H$> (HM@@@@@@@@@@@@@@@@@@@@@@@@N+_ 1_`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("wKjaJ= `_?Yt....-JWYTHYY(>?!~    `                                `  `      `  `      `      `                        ``d@@@@@@@@@@@@@@@@@@@@@@@@@@@MN, ,H@@@@@@@l `?H@@@@@@@@@@@@@@@@@@@@@@@@@@0: (>\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("1dsVY77=?^^???~~_` `  `  `         `                `   `  `                             `                             .W@@@@@@@@@@@@@@@@@@@@@@@@@@@@@H.(@@@@@@@H_`.`(@@@@@@@@@@@@@@@@@@@@@@@H@@K:(J?.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `                                            `               `        `         `                                    (M@@@@@@@@@@@@@@@@@@@@@@@@@@@H@Ml,@@@@@@@@{ ..W@@@@@@@@@@@@@@@@@@@@@@@@@@@b.~(1\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `                        `              `           `  `    `  `      `       `                      `` .H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@NJ@@@@@@@@@p.(@@@@@@@@@@@@@@@@@@@@@@@@@@@@N~..?-\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("        `        `                          `    `  `     `    `               `      `  `    `                      `,@H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@MM@HgH@@@@@@@@@@@@@@@@@@@@@@@@@@@N(__`+`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `      `  `     `  `           `                                                `                                `.H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@FG?? .-\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `         `            `       `   `     `        `   `   `   `   `  `                 `                     ` (@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@].<1(.1\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `      `          `  `                    `       `                         `       `  `                          `(@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@H@@@: 1~/:.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("           ` `                  `      `                   `   `  `      `       `  `         `                      d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#` .<(!j\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `  `       `  `      `  `      `            `    `              `      `                                    ` `.M@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   ij_J\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `           `      `                         `    `   `                      `         `  `                       (@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@HM@`  .1_n\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf(" `  ` ` ` `` `  ` `  ` ``  `  ` ` `  ` ` ` ``   `  `   `  `` `` `` `  `` `` `   ` `` `` `  `  ` `` `` `` `` `` `  ` JHHM@@@@H@@H@@H@@H@@H@@H@@H@@H@@H@@H@@H@@H@@H@@H@@@@H@@@H@@@H@@@H@@H@@H@@@@H]`  `()!\n")

	case 2:
		story2 := "        .MMMMM@                  .MMM@dMMMMY           (MMMMF  .MM%`    (MMM%                                                                          dMM`                        JMMMMM@\n" +
			"        .MMMMM@                  dMMFJMMMM'           -MMMM@  .MMf`    .MMMF                                                                          .MM#                         MMMMMMl\n" +
			"        JMMMM#`                 (MM#(MMMM'           .MMMMM@ .dMf      dMM#`                                                                          .MM#                        .MMMMM#`\n" +
			"       `JMMMM@                 .MMMldMMM^`           dMMMMP  jMC`    `(MM#'                          ``                                               ,MM#                        (MMMMMD\n" +
			"        dMMM#!                .MMM$.MM#'            (MMMM#! (MD ``   .MMMt                    ` `._` ,t                                               .MM@`                     ` MMMMM#{\n" +
			"       .MMMMD  `            ` dMM#(MM#>           `.MMMM#!`.M@` c    jMMF                  ` J. .r./! } .   `                                         ,MMb                       .MMMMMD:\n" +
			"       .MMMM{               `(MMM>dMMC          ``.HMMMM> .d#~.?`   .MM#`               `   .!},J=`   jk(...-!  `                                     .MM#`                     `.MMMMM>~\n" +
			"     ` jMMM#_         `     .MMM%.MM$`       `    dMMMMD `(M>-z     dMM>  `         ` ` `- .: ?          .vI+++<j^`                                   ,MM@                      `dMMMN#(\n" +
			"       WMMMb   `   `      ` dMMf(MMF.            .NMMM#  .MD_+`   `(MMF              .U1.,7^                ` .!                                      .MMN                       dMMMME~\n" +
			"     `.XMMM@              `.MMH!(M#~`          `.MMMM#! .M@_+:    .MM#`                1,    `               ?OVC!    `                             ` .MMN                      .MMM#H$\n" +
			"      JMMMMl              .HMM%.MM%         `  .dMMMMt  (M<+>     (MMt           ` <1Z>                          ?7!(-`                               .MMH.                     (MMMHD>\n" +
			"      jMMM#! `         `  dMMD_dMF-      `     (MMMM@``.MD<_`  ` .MM@_         `` (J~``                          (Z(,                                 .MMH_                  `  dMMMHy~\n" +
			"    `.dMMM%            ` .MMM!.M#!            .NMMM#! .M@<<`     dMM%    `          ?l        `                  ?,`                               `   MMK_                    .MMM#XC\n" +
			"     (MMM#`     `       .HMN%.MMt `           dMMMMl  (#I?`     .MM#` `            ?Y^        ...   ` `         `(nI.                                  MM#}                    .MMM@Z<\n" +
			"   ` dMMM@             `JMNf jM#`           `.MMMNF  .M@<<  `  .MM#>            ` ?=?``    `  (ds.   .` `    `   `,> `                                `dMb)                  ` JMMNPC:\n" +
			"    .MMM#_           ` .MMH:.MMl           `.dMMM#`  d#1:      dMM@             `.._.       _([hzZn.`<j+ C,`       T6.`                                dM#}                    dMMMHb~`\n" +
			"   `(MMMD     `       .WMN@ dM@       `     (MMMM% `(M@<     `.MMN}               ` .~  `   (Z0dIO/v-( ?6.>i.,;  ...i                                  dM#: `                  MMMMMS\n" +
			"    (#M#>.           `(MMK`.MM>             (MMM@  .M#6:     .dMM@       `         .7<`..` 1.WDd21o.~(-`.!1.?.z+_-l                                   `JM#<`                  .MMMNM%`\n" +
			"  ` dMMD(`        `  .HMW@ dMF   `        ``(MM#!  dM#I    ` (MMM$                 `.: x?l.?>dmssuda,..(- J(sdPd;-J' `                                 JMN<                   .MMMM#:\n" +
			"   (MMMZ$          ` Jd#f .MN!            `.MMMF `(MMK!   `.MMMM#`                 .v?jx<((udr(-dBSn/!?! (WKi/.pC?I                                   `JMMz                   JMMMM#}`\n" +
			"   dMM@j!    `      .MMX: dMF         `    dMM#! .MM#l   `.MMMMMF                 `   JCwljJdK<X-..    ` - _` Jcj.  `                  `  `            ,MMz`               `  dMMMMH!\n" +
			"   dMMDZ            dWH$`.M#!      `      .MMMF .dMM@   `.MMMM#^                      !Z<n-l vO8~-?`     j   .l  `            `  `                    `,MMZ                   MMMMMD .  `\n" +
			"  .MM#<%        `  .MNX! dMF           ` (MMMMt (MM#!  `.MMMM@`   `                    `(b.I. Z<     `  `(_ .r                                         ,MMw.                 .MMMMNF\n" +
			"  jNMr(`         ` dMH% .MM!   ``     ` (MMMM#`.MMMF ..MMMMMF            `              `0 .d+?.       _?` .>                                          .MMk_                 .MMMMMl(\n" +
			"  dMMC{      `   `(M#S  dM@ `         `.MMMM^'.dMM#(JMMMMMMF   `   `                   `.A+`?wV+.`  .<?^^!,'                                          `.MMC~               ` JMMMNM!<\n" +
			" .dM#<!          .MMN{ .MMl          `.MMMM'  (MMMMMMMMM#^~    ..(gm.`           `  `   dN?z++zwX&.``   ./                                             .MMR_               ` dMMMMN-}        `  `\n" +
			" .MN@(          `(MMZ` dM#         ` .MMMMl  .MMMMMMMMB'` ..+NMMMMM=                `..dMNc+?-jwUAvG.  .=                                              .MMb;                 dMMMMN(~    `    `\n" +
			"`jMM${           dM#: .MM%          .MMMMl  .MMMMMMM8+.&NMMMMMMMY              `  ..gMMMMMMe1..?jzwAw0dN, `                                           `.MMb}                 dMMMM#J       ` ._>\n" +
			" dNH<`       ` `-M#%  jM#`       ``.MMMM@  .dMMMMMMMMMMB^!-(MMD`          `  `` .gMMMMMMMMMMp?+?77zzwOkMNa,`                                            MM#}                MMMMMMNP     `._` .\n" +
			".WM@.           dMF` .MMN.  ` ` ..(MMMM%  `.MMMMMM#^!   .(MMY           ``  ..gMMMMMB(MMMMMMMN,4J...zwI4MMMm,``  `                                    ` MM#}                MMMMMMM]   ` _   .!\n" +
			"(MMF~          .MM:``jMM#  ` .NMMMMMMM%` .(MMMMF`      .MMD`        `  ..+MMMMMMMB=(gMMMMMMMMMN,(G????Z0VMMMMN,. `  `                                   MMNr             `  MMMMN#T%`   ~  .!\n" +
			"dMM%`      `   (M@` .MMM#` .MMMMMMBBMNgMM#=         `.dM@`       ``..gMMM#^=! ..gMMMMMMMMMMMMMMMN,?<..JwWvMMMMMMNg,.  `  `                              MMMD           `  .JMMMMM]   ` -._`\n" +
			"HMM:        ` .MMl  jMMM$.MMMM#UQgMMM#^!`        `  .MM^        `.MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMm,<+ZY<1,MMMMMMMMMm,                                  dMMb`         ` .JMMMMMM#!    .`\n" +
			"WM#`   `    .MMM#  .MMMMMMMMMNMMMM^!          ` ` ..dMF          MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMm,i.._3(MMMMMMNmTMMb                              ` dMM#       `   .MMMMMt                      `  `\n" +
			"MMF`     `.dMMM@!` JMMMMMMMMMM9!          `` ..+MMMMM#`         .MfMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMp_-_? MMMMMMMMMMMN`                               JMMN.    `  `.jMMMMY`                    ` ` ..a\n" +
			"MM]    ` .MMM#'`  ` MMMMM#^?`          ``..JMMMMMMM#=`          ,M](MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN.(,~JMMMMMMMMMM#  `                          `  ,MMM[      `.MMMMB!     ` `           ` ` .JQNMM\n" +
			"MM}      dMM@`  .J&M^MMM%            ..gMMMMMMMMMMNgZ       `  `.MNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,(),MMMMMMMMMMF                                ,MMM]     `.MMMMD        ./!     `  ` ..dQNMMMMM\n" +
			"MH`     .MMMi(NMMB^   ?!     `` ..JNMMMMMMMMMMMMM#^!             MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN,>.MMMMMMMMMMl `                              .MMMb  `  .MMMMt     `  -_  `  ` ` ..dWMMMMMB^`\n" +
			"M#   ` .dMMMMM9!         ``..JgMMMMMMMMMMMMM#^^^                 dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMb MMMMMMMMM#                       `  `    ` .MMM#   `.MMMM%          `    `..(WWMMMMB=`\n" +
			"MF    `dMMM9!`     `  ..JgMMMMMMMMMMMMM#^^               `    ` .(MM9TYTHMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMmdMMMMMMMM#                         `       (MMMN  `.MMMMD              ..(dHHY?U7=`\n" +
			"H%     ?B=      ` .gNMMMMMMMMMMMMMM9=!           ` .`          .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMF                     ` <??<.    ` ?MMN. (MMMMt  `. `     ` .(QNHY!`\n" +
			"W:     `.(gNg,.(gMMMMMMMMMMMMM#^^ `              .!_           `.MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM]                         ~.(`      .MMN.MMMMF   ` ! `  .-JQWBY!\n" +
			"K`  `` .MMMMMMMMMMMMMMMMMMM9=`                                 `.MMMMMM@1JMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM%                             `      ?MMMMMMM$ `     .gkQHHV=`\n" +
			"#`    .MNMMMMMMMMMMMB^! _                     `          `   ` .MMMM#^.gMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN!                                     MMMMMMF   ` .+MMMMY^\n" +
			"D    .MMM$ MMMM#^!                                    `        dMM@~-MM5MMMMMMMHMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#                                       dMMM@   .gMMMM9^\n" +
			"%    MMMM! /B^                ` .                             .MMNMMMY(MMMMMMMNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMF                                       ,MM#~.gMNMM9^                      ` .\n" +
			": ` .MMMF `                     ~.~-`                       `.MMMMMD(dMMMMMMMNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM% `                                      TMNMMMM#=                       ` ,Z`\n" +
			"`   dMMM!                         `_            `            (MMMM3jMMMMMMMMFdMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM:                            `    ` `` .gMMMMMY` `                     `.,!.!   `\n" +
			"   .MMMF                               .   `              ` .MMM@(dMMMMMMMM#`.MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#                               `  ` .+MMMMMD!                       ` ./~..!`   `..\n" +
			"  `(MMM:           `                   `(``           `    `.MM#jMMMMMMMMMM>  (MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMF                         `       .JMMMMMB^                        ` .<_-(+:  .(gMMM\n" +
			"   MMM@ `     `         `                                   dMMNMMMMMMMMMMF    dMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMb                            ` .JMMMMM#=                        `` .<__+zz>  .MMMMMM\n" +
			"  .MMMl                                                    .MMMMMMMMMMMMM@     .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN                          ` jMMMMMM^`                        ``.JMC+1zzwIJgMMMMMMMM\n" +
			"  (MM#  +.`                                     `        ` jMMMMMMMMMMMM#!      JMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN.                      ` `.dMMMMY!                       ` ` .kMMDdkwWWSMMMMMMMMMMM\n" +
			"  dMMt  }(.     `     `   `                               .MMMMMMMMMMMMMt        MMMMMMMMMMMMMMMMMMMMMMMMMMMMM<MMMMMM;                        `.MMMF                       `   .JMMMMHWHMHHSMM#^!`?TXkT^\n" +
			" .MM#`  ( <.`                                         `  .MMMMMMMMMMMMM#        `(MMMMMMMMMMMMMMMMMMMMMMMMMMM#.WMMMMMb `                       MMMMl                   `   ` .dMMMMMMNMMNM$`\n" +
			"`jMMF    { (                          `      `          `dMMMMMMMMMMMMM{          dMMMMMMMMMMMMMMMMMMMMMMMMMM@xzMMMMMN.                      ``.MMN`               ` ``..J1JMMMMMMMMMMMMM^\n" +
			".MMM:    _- ~. `                   `                     MMMMMMM9udMMMM}          (MMMMMMMMMMMMMMMMMMMMMMMMM@.IWdMMMMMb                         MM@             `` ..gMBQMMMMMMMM9HMMMM#'       `\n" +
			".MM#       _<<.`   `    `                            `  (MMMMY~(MMMMMMD`          .MMMMMMMMMMMMMMMMMMMMMMMMM!J_(bdMMMMN,`                      .MM>        `  ..JNM#5jdMMMMMMMB:_~~~.(?`\n" +
			"jMM]                                     `             dMMB=-.MMBMMMMF         `   MMMMMMMMMMMMMMMMMMMMMMMMF,^-~j/MMMMMN,`                  ` .MMt`   `  ..(gMMMMYqgMMMMMMMMY~~~~~-J!\n" +
			"MM#!     `                                      `     `JMMm+MMD`dMMMM[   `         MMMMMMMMMMMMMMMMMMMMMMM#(>...(X,MMMMMN, ` `                dMt`  ..JgMMMMVY<j+MMMMMMMM8:~~~~~(?!\n"

		fmt.Printf("%s", story2)
		time.Sleep(1200 * time.Millisecond)

		fmt.Printf("MMl                     `                    `        `.MMMM#' (MMMMM]             dMMMMMMMMMMMMMMMMMMMMMM:J~__~(yh/MMMMMNx             `  ` jM5.JgMMMMMBSz!.+MMMMMMMMB=_~~~~(?!`                  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("Mt                                 `                   (MMM#~.(MMMMMM@`           .MMMMMMMMMMMMMMMMMMMMMMF.l .._~jVh,MMMMMMp  `           `..MMMMMMHHs77~.JMMMMMMMMMY<:::~((!                   `    `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("@                  `    ,`               `            .MMMF.MMMMMMMMM>         ` .dMMMMMMMMMMMMMMMMMMMMM#!,?!``~(vGd;(MMMMMMN,            ,MMMMMM9^!  .JMMMMMMMMM5?<<<;++!`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       `        `       l                            `.MMM(MMMMMMMMMF        ` `.MMMMMMMMMMMMMMMMMMMMMMMF m(+ggNMMMMN..MMMMMMN.          ``?UY^  ` .JMMMMMMMMM81O11+zv?`                               `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                                      `             `.MMM@jMMMMMMMM#            dMMMMMMMMMMMMMMMMMMMMMM#!.MMMMMMMMMMML`,MMMMMML`             `  .JMMMMMMMMMBOOOzzzZ?`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    `        `                     `            `  ` .MMMgMMMMMMMMMl          ` dMMMMMMMMMMMMMMMMMMMMMMF (MMMMMMMMMMH8l dMMMMMN,`      `     .+MMMMMMMMMBwwZOOrC?``\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                          `     `  `.(`              dMMMMMMMMMMMM#      `    `.MMMMMMMMMMMMMMMMMMMMMM#` ZTYC<_ ..~  (W[.MMMMMMN,  `    `..gMMMMMMMMMHXXXuXwZ^`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                                   (<! `     `      .MMMMMMMMMMMMM#           .MMMMMMMMMMMMMMMMMMMMMMM@ .!`??!+J>`.S&d8d;(MMMMMMMp   `..gMMMMMMMMMHWXyXuXT^`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                   `    `                           (MMMMMMMMMMMMM@         `.dMMMMMMMMMMMMMMMMMMMMMM@  jWHC-.J(k,(C_?U9W,?MMMMMMMb.(MMMMMMMMMMWHWHVWWY!                                   ` `  ...`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("             `                                   ` .MMMMMMMMMMMMMM:        ` gMMMMMMMMMMMMMMMMMMMMMM#` .wsxvXy>.XFdk<- .TN,-MMNMMQMMMMMMMMMMHHHHWWB^`                                ` ` `..!~`.~`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    `  `        `                                  .MMMMMMMMMMMMM#          jMMMMMMMMMMMMMMMMMMMMMM#'  g(gMMMt: zHGXHJ_. (N,.QNMMMMMMMMMMH@HHHHB=`                                  `   .!   .~\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                                             `   ` dMMMMMMMMMMMMMb      `  jMMMMMMMMMMMMMMMMMMMMMMM' `.MMMMMM!` .8XMkH-?XXQMMMMMMMMMMMHMMHNM^!                                    `   .!  ..(!    `` `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                      `   `        `     `        .MMMMMMMMMMMMMMF        jMMMMMMMMMMMMMMMMMMMMMMMl  .MMMMMMMNJ,  (SUWgNMMMMMMMMMNMMMMMMB=`                                         ,~..?!      `.._~?`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("             `  `                                 (MMMMMMMMMMMMMMt `  ``.dMMMMMMMMMMMMMMMMMMMMMMM'  .MMMMMMMMMMMMHigMMMMMMMMMMMMMMNMMHNMMMm,                                       _   `        __!  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("          ` ` ...v`                             ` dMMMMMMMMMMMMMM:      dMMMMMMMMMMMMMMMMMMMMMMM' `.MMMMMMMMMMMNMMMMMMMMMMNMMMMM#^CMMMMMMMMMm,                             `  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     ` `` ./!`.!                               ` .MMMMMMMMMMMMMMN`     .MMMMMMMMMMMMMMMMMMMMMM#!  .MMMMMMMMNNMMMMMMMMMMMMMMMMMNm.  (MMMMMMMMMMN, ` `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       ./`  `,`         `          `     `    ` .MM#dMMMMMMMMMMM#      dMMMMMMMMMMMMMMMMMMMMMF   .MMMMMNNMMMMMMMMMMMMMMMMMNMMMMMN,` _?MMMMMMMMMMNJ.`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    ``.!  ` .`                               ` .MM#(MMMMMMMMMMMMN.    .MMMMMMMMMMMMMMMMMMMMMF   (MMMNMMMMMMMMMMMMMMMMMNMMMMMMMMMMN,` _<?MMMMMMMMMMNa,  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     ,`  `./`         `                      `.MMM>dMMMMMMMMMMMMM]  `.dMMMMMMMMMMMMMMMMMMMM=  .jgMMMMMMMMMMMMMMMMNNMMMMMMMMMMMMMMMMp`  <_(TMMMMMMMMMMN,. `  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `.```.~      `                  `         .NMM@.MMMMMMMMMMMMMMb   jMMMMMMMMMMMMMMMMMMM#~.+MMMMMMMMMNMMMMMMNNMMMMMMMMMMMMMMMMMMMMMR.  (<..TMMMMMMMMMMMNJ.`  `                                   `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    :`./``                `              ` `.MMNadMMMMMM@=?~__.TMF` (MMMMMMMMMMMMMMMMMMMNMMMMMMMMMMMMMMMMNNMMMMMMMMMMMMMMMMMMMMMMMMMMN.  _1- .TMMMMMMMMMMMMNJ..  ` `                               `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `.>!                                      MMMMMMMMMMMD.(z((KQMt  .MMMMMMMMMMMMMMMMNMMMMMMMMNMMMMMMNQMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMKi   ?<- ,HMMMMMMMMMMMMMMMNJ,  ` `                        ` .~~ ._\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `    `         `  `     `     .TMMm-(dX&,?`  JJM9! `.MMMMMMMMMMMMMNMMMMMMMMMMMMMMMQgMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#<1.   ~:-.?MMMMMMMMMMMMMMMMMMN,.  `                       I!!!!\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                        `    .,`               ?TMMMNmr    ?MMN&,.MMMMMMMMMNgMMMMMMMNMMMMMMHgMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#MMMMMM#J6;vG,  ~<-_?MMMMMMMMMMMMMMMMMMMm,.`  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                          `.de7!   `               `_T;` dQXmJ,7HMMMMMMHgMMMMMMMMMMMMB^jgMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM#MMMNJMMMMMMK1Sz?+1..  ~_<?TMMMMMMMMMMMMMMMMMMMm,   `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                          .8f                        .0_.dHHMMBH@MMmgMMMMMMMMMMMMMK`  (MMMMMMMMMMMMMMMMMMM@dMMMMMMMMMMMMMMMNdMMM;HMMMMMMrzSZJ+i?G,  ``__TMMMMMMMMMMMMMMMMMMMNa.`  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `       ` ` (d~                       `zz<<XNIHM1MNMMMMMMMMMMMMNNMND  .dMMMMMMMMMMMMMMBMMMMMlJMMMMMM@MMMMMMMMM/MMM],MMMMMMN(zXI(C<1(T+.     .TMMMMMMMMMMMMMMMMMMMN, `             `  ```\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `               `     .X}             `          .tz(>d#1KMdMMMMMMMMNNMMMMMMM@  (MMMMMMMMMMMMMMM5WMMMM#>dMMMMMM$?MMMMMMMML?MMb dMMHMMM;vvO111z+(-vG,       7MMMMMMMMMMMMMMMMMMN&.` `         ` (-~?\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  `,C`                   ,d)                 `  `  .(_z(>dNjMMd0ZM9NMMMMMMMMMM#! .dMMMMMMMMMMMMMMM$xdMMMMv(HMMMMMMt(dMMMMMMMN.HMN.(MMPMMM@(v+l1z<11?-1(7&,`     .TMMMMMMMMMMMMMMMMMNe.  `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   v                   ` (Hb   `   `  `     `  ..JcH$(>.;jMdWHm9IHHMMMMMMMMMMF  (MMMMMMMMMMMMMMMMDJjMMMMD!.MMMMMMM{.dMMMMMMMM;,MM;.MMbdMMN(zC1:_~<<?(1?_-?Tu,       TMMMMMMMMMMMMMMMMMa. `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("  .r         `            ?Vb        `  ...JQHjHiHmN$_}.:_?jdDHHQM#MMMMMMMMM@ .MMMMMMMMMMMMMMMMMD,UMMMM#> dMMMMMMN_.MMMMMMMMMb dMb MM#(MMMb(111<. _(((v<1<~<<?4-..`    7MMMMMMMMMMMMMMMMe.`     ` .\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                `          (HWG,.` .JXH9MhXNdNMHHY^Ti.,(3,`.KMNmaxWMMMMMMMM'.dMMMMMMMMMMMMMMMMMD.DjMMMMD .MMMMMMM#  dMMMMMMMMN..MN.dMN_dMMbj+1z((~. -_~<+1/11/<<(??T(.`   7HMMMMMMMMMMMMMNe.``\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       `                  `  `?7WHY7NuHHMB^=!       ?7Gc-W+.(_dNghdMMMMMM^ .MMMMMMMMMMMMMMMMMM# dkMMMMN! dMMMMMMMX_ JMMMMMMMMM| dMp,MM].MMNv1Gzl~1-.`  `_<(-~?><<-1.1(T-   ` ?TMMMMMMMMMMMMN&  `     `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                      `           _7!                  (&dHa+dMMMMMMMM9^ .(MMMMMMMMMMMMMMMMMMM#(NMMMMM$`(MMMMMMMM#  ,MMMMMMMMMb .MN(MMb.MMMIi>jvt  ~.   ` `_<(+11<<?,<<-?&      TMMMMMMMMMMMMN,   `\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      `      `     `                                  ` .dMMMMMMMMMMY`  .MMMMMMMMMMMMMM#dMMMMM$jMMMMM#!.MMMMMMMMM#. .MMMMMMMMMM; dMZMMN.MMMb+11<+l_.         ~.><+11(<<+-_?, `  ` 7MMMMMMMMMMMMN,\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    ` .                                                .dMMMMMMMM#=`   .MMMMMMMMMMMMMMEvdMMMMF.MMMMMMD.(MMMMMMMMM@_  MMMMMMMMMMN.JM[MMM[MMM#(+<1;1v(.        `  ~<11?<=/?J-<1.     .WMMMMMMMMMMMM8+.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("      ?-  `                                           .dMMMMMMM@!    .dMMMMMMMMMMMMMNE>-MMMMF JMMMMMM: MMMMMMMMMMb~  dMMMMMMMMMM),MNdMMNdMMN:1<_<_?JZi.           _?+i11-(z(1?i      ?MMMMMMMMN>   _i.`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("     `..~.`     `    `  `                           `.dMMMMMM^      .MMMMMMMMMMMMMMM@>`.MMM#  MMMMMM#`.MMMMMMMMMMb_ `dMMMMMMMMMML,MMMMMMMMMMr<x~` _(+1?>.```         <?J?I+O<O+1,`  .,.WMMMMMMMN,\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("   `   _. !-.``                                      dMMMM#^     `.MMMMMMMMMMMMMMMM#=  dMMM# (MMMMMMD JMMMMMMMMMMb`  ,MMMMMMMMMMN,MMMMMMMMMMb((i.  -(<1<z<. .`        _(O+O+1Jvo?l   I1.?MMMMMMMMp\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    `    - ` ~!,  `                      `       `  jMMM@!      .gMMMMMMMMMMMMMMMMMF  .MMMM#`(MMMMMMl dNdMMMMMMMM@.  .MMMMMMMMMMMyMMMMMMMMMMN+((+.   __<_1<_~_          (1(?1?&+IJI. (I?G,WMMMMMMMm\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("       `  _.   .-`                           `   `.uMMD`  `  .JMMMMMMMMMMMMMMMMMMM#`  dMMMMF.MMMMMMM: dMMMMMMMMMMN_ `JMMMMMMMMMMMNMMMMMMMMMMMr11(1.  (-__ _<__`          (+__?(Ozuc1i.x011i?MMMMMMMb\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("          ` ~-....        `        `           ` .MMY   ` .JMMMMMMMMMMMMMMMMMMMMMMt  .MMMMM>dMMMMMMM!.MMMMMMMMMMMN>  dMMMMMMMMMMMMMMMMMMMMMMMb(Ii     <_~_  `__  `        (<i. _1(vn1z?1zIzzJMMMMMMMN,`\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("              `                 `           `  .MMMF   .uMMMMMMMMMMMMMMMMMMMMMMMNN,  dMMMM#.MMMMMMM# (MMMMMMMMMMMM; .MMMMMMMMMMMMMMMMMMMMMMMMN<+11-`  .__~_    _           <<>   (<<WJzXJwzv+/WMMMMMMMp.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                      `                  `  ` dMMMM: .2(MMMMMMMMMMMMMMMMMMMMMMMMMM@`.MMMMMNMMMMMMMMN,dMMMMMMMMMMMMr` dMMMMMMMMMMMMMMMMMMMMMMMMr<(zz-` `_:.`              ` _<<_    ~~C+1JvvXoj,?MMMMMMMN\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                   `                  `       ,MMM@.WU1MMMMMMMMMMMMMMMMMMMMMMMMMM@ .MMMMMMMMMMMMMMMMtMMMMMMMMMMMMM#..MMMMMMMMMMMMMMMMMMMMMMMMM@??(<z_. .<~                 _<<_.     _1Ix1+dmsOijdMMMN?^\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                                               (M#jSI1dMMMMMMMMMMMMMMMMMMMMMMMMMM[.MMMMMMMMMMMMMMMMMgMMMMMMMMMMMMMM|(MMMMMMMMMMMMMMMMMMMMMMMMM#((+<-+((<<__                 ._~_       1?zI?TMN/<x7MMMN.\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("    `                     `        `            MBXXtdMMMMMMMMMMMMMMMMMMMMMMMMMMMnMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNMMMMMMMMMMMMMMMMMMMMMMMMMMN<<+<+_i<<<<_   `               <_       `_(1?/?MMm.g.dMMN\n")
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("                                               `dmIuMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM6<<<<1_(. ___                              ~<__?MMMMM#^MM\n")
	}

	text3 := "																																																																																																		 \n" +
		"																																																																																														 				\n" +
		"        .+XXT<`                                             `                         `                       `                  `            `        `        `                         ?7WXA.\n" +
		"    ` .dY(X{`   `            ` ` .JuwwwwwwwwwwwwwwwwwwwwwwA&-.` `..uwwwwwwww>`  ` `..&wwwwwwwwwwwwwwwww2`  .(uwwwwwwwwwwwww2 ` .JuwwwwwwwwwwwwA&..`  .Juwwwwwwww:.JwwwwwwwwZ `             `(X/Ck- `\n" +
		"  `   J0 .Wl.`                `.dV1(((((((((((((((((((((((-J7yo JXYi((((,.df``  ` JWYi((((((((((-((,.d0``.dV=+(((((((((--JX!.+XY1((((((((((((--?4k,.dV=(((((,(ykXY1(((((-jX!  `         ` ` .0~ d$`\n" +
		"    ``?W,` 7UA-..`` `     ` ` .Xf_7^^^^^^^T@@@@@M^^^^TH@@@@Mr(WwV>j@@@@M%JX!  `  (y}j@@@@@Y^^^^^^^>(dV! .X$(H@MH@@@@@Y=(dV>.XY(d@@@@M9^^^TM@@@@H;dkdf(W@@@@#<Xyyr(@@@@@D.yC`    `   `   .(dV=` .y%\n" +
		" `     ?Uk-.  _74k.    `    ` JUUUUUUUUWW>d@@@@M%JVVVlJ@@@@M%JffC(@@@@@F(f%    `.Wf.@@@@@D(VUUUUUUV7!  (Vt(@@M=d@@@@M~Xf^`.W$.@@@@@@(XVVX(W@@@@#<XfW:d@@@@HlJVVS(H@@@H#(Xf   `       .dY=!  .(dV^   `\n" +
		"   `      ?TW;`  .Wl  `  ` `  ..      .Wt(@@@H@F.fff$(@@@H@D.fpf(H@@H@#(f$   ` .df(H@@H@#(Xf         .JW>J@H#<.H@@@H#.f@  dS(W@@@H#<XffWld@@@H@%jff%(@@@@HF(ffW:d@H@@MldW!. `  ` ` `(S!   (XY=`\n" +
		"       `   `JR    dD`    .dWpf^  `  `.Xf(H@@@@#(Wpp0(H@@@@#(ppW<d@@H@M>dk`  `  (W>dM@H@MldpA+++++++..dK<d@@@(f.@@@@@F(p: (W>d@@@HM%JppW3(@@@H@P.ppf(@@@H@@(Wpp%(@@@@H$(W}?UpWA..`  `dR  ` WP`     `\n" +
		"        ` `.Kl  .JK! ` .dY_df`     ` JW~d@@@HMlJ+z+(H@@@#=(Xpb%(@@@@@t(W>     .W%(@@@@@b.+1+1++?dW}.X9(d@@D(p$J@H@@@@dK_.Wt(@@@H@P(pbpR_^^^^^^(ppK<d@H@@MlJ1zi(H@@@@@(b$   4R.Ck,   .Wn.  ?H;  `   `\n" +
		"    `  ` .j9!  .df`  ` dD .W]     ``.b%j@@@@@YYYYYM@@@Ha,Wbbk$(@@@@@@(b%  ` `.Xf.@@H@@#YYYYYYY1Jb3.Wf(H@M$.77>d@@@HM:pD.Wf(@@@@@#(WkbK=7777777Wbk%j@@@@MYYYYYW@@@H@M<Xf`   ,K_ dR     4R. `(4n. ``   `\n" +
		"      `.J9=    jH`     db  dR.`   `.Wf(@@H@@D-WWkD(@@@@@FJkkK(W@HH@#<XK`     dK<H@@@@M<dWpppbb9Y!(k=(@@MHHHHHH@@@@@#(k$dK<d@H@@M:dkkk%(@@@@@t(Hk$.@@@@HF(WWWld@@@@M%jW!    dD  dK     .Wr    ?Wn,\n" +
		"  `  `.H'    ` WD  `   .H) .Wh. ``.dK(W@@@@#<XqqH<W@@@HM<XqH>d@@@@M%JH!  `  (H@d@@@H@tjH:  ` ` .jH3(@@#<dkkkK.@@H@@D(qkq%j@@@@@tJqqqD(M@H@@@.qqK(H@@@@#(WqH$(@@H@@F.q%  `.jK` .H>   `  JR.`    ?Ho` `\n" +
		"  `  .W%       Wb   ` `.Hl   TH+.`(Hld@@@@M3dmqHlj@@@@M%jqqt(@@@@@D(q%    `.H$(@@@H@D.m%      .d#<d@@@(WqmmqP(@@@H@@dqmP.@@@@HD(mqqK(W@@@H#<WqH>d@@@HMldqqK(H@@@H#(HD   .XY`  (H<  `   JK`      (H_\n" +
		"`.J  .Ho `` `  (Wm....+W=` ` `.UHQm$(@@@H@F(mmg$(@@H@@F(mmK(@@@H@#(WHJJJJJJXK(H@@@M#(WH+(JJJJJWK(d@@Y(Hgmmgm@d@H@@H>XmK(H@H@@#(WmmHld@@@@M@dHmt(@@H@@$(gmH:d@@@@M>dHa-+WH&((,  ?Hm-...d9!  `  ` (H`  j.`\n" +
		".X:  `(WmJ...`  ` ?777`      .WBT71(H@@@@#-777iJ@@@HMD(HgH<dH@@@Ml?17777777H>d@@@HMl?77777777WD(H@M%(ggggggg:W@@@@N(17>J@@@H@h(771(d@@H@H3(ggD(@@@@@@(Hgg%J@@@@@D(CC717777XH!   `-?777`` ` ....d#^   .H,\n" +
		"JN_    `_7TYWHm,           `.H%JWHWWWWWWWWWWWWWWWY9=(d@@@%?WWWWWWWWWWWWYYid%?BWWWWWWWWWWBYYid$(HYYjdg@g@g@@@e7BHWWWWWWWWWWWWWWWWWWWWWYY1(Wg@H~TWWWBT(WgHD(BWWWWWWWWWWWWY=(Hl            .gHHYY^!`     Hb\n" +
		"HHL`          zN-          .HMbWkkkkkkkkkkkkHkkkkHH@@@@@@kkHkkkkkkkkkkkHH@@HkHkkkkkkkWkkkk@@HHHkH@@@@@@@@@@@@HHHkkHkkkkkkkkkkkkkkHkkkH@@@@@@HkkkkkH@@@@@HHHkkkkkkkkkkkkH@H+..        ` .HY        `` (@M\n" +
		"HKWm,      `  .H! `          .d@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@M@7777777774Y7U74#7777B^777H@@@@@@@@@M#YY^^^^^7H@H@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@M%  `    `  .Mp    `  ` .gBWH\n" +
		"JN-?TNaJ....(gB'           `.H@@@@H@@H@@H@@@@@@H@@@H@@@@@@H@@H@@@H@@H@@@@H@HHMHHHH#`` .Z J{.H%`  (l   (@H@@H@@HH@MaJ.   +md@H@@@@@@@@H@@@@@@H@@@H@@@@@@@@@@@@@@@H@@@@H@@H@HD.           (WHJ....(gH9^.HF\n" +
		" WN.   ?7^^=!   `  `  `    .MMHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH@HHHHHM>`  jHHHHMD....P` `.Ml.......Wb...` `...JHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH@HHHD   `      ``  ?7^^7!` `.d@`\n" +
		"  TNx.`  `        `  `  `` `.MM@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@MHHHHHt   .#HHHHHHHH##`  .HH#H#####HH#M@   .M#HHHH@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@M@`       `              -MD`\n" +
		"  `(TNm.. `    ` ``  ......(dNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM######Ng++&MNN#####NNMm++&dNNNNNN#NNNNNMa++&dNNNNNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNNJ(......  `     `   ..gMB^  `\n" +
		"    ` ?THMMMMMMMMNMMMMHY^^^^^77777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777^^^^^HMMMMMMMMMMNMMH9=\n"
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("%s", text3)
}
