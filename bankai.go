package main

import (
	"fmt"
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

	text2 := "ltOOwV_dOltttttltllltlllllldf`  J0tttttlltltuV= .HSttlttttrtttrttrrttttttttttttttttttrtrtttttttttttrttttrttttttrttttttttttttttttttttlltltltltlltlllllllllllltltllllttlllllllllllllllllllllllltOrrrvzrrOt\n" +
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

	fmt.Printf("%s", text1)
	time.Sleep(1500 * time.Millisecond)
	fmt.Printf("%s", text2)
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

	text3 :=
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
