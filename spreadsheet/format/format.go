//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package format provides support for parsing and evaluating
// spreadsheetml/Excel number formats.
//
// Internally spreadsheets store numbers and dates values as a text
// representation of a floating point number (e.g. 1.2345).  This number is then
// displayed in Excel or another spreadsheet viewer differently depending on the
// number fornat of the cell style applied to the cell.
//
// As an example, the same value of 1.2345 can be displayed as:
// - "1" with format "0"
// - "1.2" with format "0.0"
// - "1.23" with format "0.00"
// - "1.235" with format "0.000"
// - "123%" with format "0%"
// - "1 23/100" with fornat "0 0/100"
// - "1.23E+00" with format "0.00E+00"
// - "29:37:41s" with format `[h]:mm:ss"s"`
package format ;import (_e "bytes";_b "fmt";_ee "github.com/ScorpioRen/copy-unioffice";_c "io";_f "math";_df "strconv";_g "strings";_d "time";);func (_ab FmtType )String ()string {if _ab >=FmtType (len (_gg )-1){return _b .Sprintf ("F\u006d\u0074\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029",_ab );};return _cf [_gg [_ab ]:_gg [_ab +1]];};func Parse (s string )[]Format {_ada :=Lexer {};_ada .Lex (_g .NewReader (s ));_ada ._cedb =append (_ada ._cedb ,_ada ._bada );return _ada ._cedb ;};

// AddToken adds a format token to the format.
func (_abd *Format )AddToken (t FmtType ,l []byte ){if _abd ._dc {_abd ._dc =false ;return ;};switch t {case FmtTypeDecimal :_abd ._bd =true ;case FmtTypeUnderscore :_abd ._dc =true ;case FmtTypeText :_abd .Whole =append (_abd .Whole ,Token {Type :t });case FmtTypeDate ,FmtTypeTime :_abd .Whole =append (_abd .Whole ,Token {Type :t ,DateTime :string (l )});case FmtTypePercent :_abd ._gc =true ;t =FmtTypeLiteral ;l =[]byte {'%'};fallthrough;case FmtTypeDigitOpt :fallthrough;case FmtTypeLiteral ,FmtTypeDigit ,FmtTypeDollar ,FmtTypeComma :if l ==nil {l =[]byte {0};};for _ ,_gd :=range l {if _abd .IsExponential {_abd .Exponent =append (_abd .Exponent ,Token {Type :t ,Literal :_gd });}else if !_abd ._bd {_abd .Whole =append (_abd .Whole ,Token {Type :t ,Literal :_gd });}else {_abd .Fractional =append (_abd .Fractional ,Token {Type :t ,Literal :_gd });};};case FmtTypeDigitOptThousands :_abd ._db =true ;case FmtTypeFraction :_ba :=_g .Split (string (l ),"\u002f");if len (_ba )==2{_abd ._fg =true ;_abd ._ce ,_ =_df .ParseInt (_ba [1],10,64);for _ ,_cfa :=range _ba [1]{if _cfa =='?'||_cfa =='0'{_abd ._gb ++;};};};default:_ee .Log ("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0070\u0068\u0020\u0074y\u0070\u0065\u0020\u0069\u006e\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0076",t );};};const _cec int =34;const (FmtTypeLiteral FmtType =iota ;FmtTypeDigit ;FmtTypeDigitOpt ;FmtTypeComma ;FmtTypeDecimal ;FmtTypePercent ;FmtTypeDollar ;FmtTypeDigitOptThousands ;FmtTypeUnderscore ;FmtTypeDate ;FmtTypeTime ;FmtTypeFraction ;FmtTypeText ;);func _bgb (_acf float64 )string {_bf :=_df .FormatFloat (_acf ,'E',-1,64);_adc :=_df .FormatFloat (_acf ,'E',5,64);if len (_bf )< len (_adc ){return _df .FormatFloat (_acf ,'E',2,64);};return _adc ;};const _cb =1e11;func IsNumber (data string )(_fdd bool ){_agb ,_gbfd ,_aafa :=0,0,len (data );_fbb :=len (data );_gfc ,_dca ,_dff :=0,0,0;_ =_dca ;_ =_dff ;_ =_gfc ;{_agb =_gbee ;_gfc =0;_dca =0;_dff =0;};{if _gbfd ==_aafa {goto _ggbe ;};switch _agb {case 0:goto _cfc ;case 1:goto _gbb ;case 2:goto _gab ;case 3:goto _aca ;case 4:goto _gae ;case 5:goto _afe ;case 6:goto _bbcd ;case 7:goto _egg ;};goto _gga ;_fca :_dca =_gbfd ;_gbfd --;{_fdd =false ;};goto _ffd ;_fee :_dca =_gbfd ;_gbfd --;{_fdd =_dca ==len (data );};goto _ffd ;_fgg :_dca =_gbfd ;_gbfd --;{_fdd =_dca ==len (data );};goto _ffd ;_ggf :switch _dff {case 2:{_gbfd =(_dca )-1;_fdd =_dca ==len (data );};case 3:{_gbfd =(_dca )-1;_fdd =false ;};};goto _ffd ;_ffd :_gfc =0;if _gbfd ++;_gbfd ==_aafa {goto _ccg ;};_cfc :_gfc =_gbfd ;switch data [_gbfd ]{case 43:goto _feed ;case 45:goto _feed ;};if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _fbbc ;};goto _aeb ;_aeb :if _gbfd ++;_gbfd ==_aafa {goto _dag ;};_gbb :goto _aeb ;_feed :if _gbfd ++;_gbfd ==_aafa {goto _fege ;};_gab :if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _fbbc ;};goto _aeb ;_fbbc :if _gbfd ++;_gbfd ==_aafa {goto _gabd ;};_aca :if data [_gbfd ]==46{goto _eaaf ;};if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _fbbc ;};goto _aeb ;_eaaf :if _gbfd ++;_gbfd ==_aafa {goto _dfbb ;};_gae :if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _aec ;};goto _aeb ;_aec :if _gbfd ++;_gbfd ==_aafa {goto _fea ;};_afe :if data [_gbfd ]==69{goto _ccf ;};if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _aec ;};goto _aeb ;_ccf :if _gbfd ++;_gbfd ==_aafa {goto _ebb ;};_bbcd :switch data [_gbfd ]{case 43:goto _ffb ;case 45:goto _ffb ;};goto _aeb ;_ffb :_dca =_gbfd +1;_dff =3;goto _dfb ;_dge :_dca =_gbfd +1;_dff =2;goto _dfb ;_dfb :if _gbfd ++;_gbfd ==_aafa {goto _bfc ;};_egg :if 48<=data [_gbfd ]&&data [_gbfd ]<=57{goto _dge ;};goto _aeb ;_gga :_ccg :_agb =0;goto _ggbe ;_dag :_agb =1;goto _ggbe ;_fege :_agb =2;goto _ggbe ;_gabd :_agb =3;goto _ggbe ;_dfbb :_agb =4;goto _ggbe ;_fea :_agb =5;goto _ggbe ;_ebb :_agb =6;goto _ggbe ;_bfc :_agb =7;goto _ggbe ;_ggbe :{};if _gbfd ==_fbb {switch _agb {case 1:goto _fca ;case 2:goto _fca ;case 3:goto _fee ;case 4:goto _fca ;case 5:goto _fgg ;case 6:goto _fca ;case 7:goto _ggf ;};};};if _agb ==_eaf {return false ;};return ;};func _cce (_aaf _d .Time ,_acff float64 ,_cg string )[]byte {_dbd :=[]byte {};_adb :=0;for _agc :=0;_agc < len (_cg );_agc ++{var _fag string ;if _cg [_agc ]==':'{_fag =string (_cg [_adb :_agc ]);_adb =_agc +1;}else if _agc ==len (_cg )-1{_fag =string (_cg [_adb :_agc +1]);}else {continue ;};switch _fag {case "\u0064":_dbd =_aaf .AppendFormat (_dbd ,"\u0032");case "\u0068":_dbd =_aaf .AppendFormat (_dbd ,"\u0033");case "\u0068\u0068":_dbd =_aaf .AppendFormat (_dbd ,"\u0031\u0035");case "\u006d":_dbd =_aaf .AppendFormat (_dbd ,"\u0034");case "\u006d\u006d":_dbd =_aaf .AppendFormat (_dbd ,"\u0030\u0034");case "\u0073":_dbd =_aaf .Round (_d .Second ).AppendFormat (_dbd ,"\u0035");case "\u0073\u002e\u0030":_dbd =_aaf .Round (_d .Second /10).AppendFormat (_dbd ,"\u0035\u002e\u0030");case "\u0073\u002e\u0030\u0030":_dbd =_aaf .Round (_d .Second /100).AppendFormat (_dbd ,"\u0035\u002e\u0030\u0030");case "\u0073\u002e\u00300\u0030":_dbd =_aaf .Round (_d .Second /1000).AppendFormat (_dbd ,"\u0035\u002e\u00300\u0030");case "\u0073\u0073":_dbd =_aaf .Round (_d .Second ).AppendFormat (_dbd ,"\u0030\u0035");case "\u0073\u0073\u002e\u0030":_dbd =_aaf .Round (_d .Second /10).AppendFormat (_dbd ,"\u0030\u0035\u002e\u0030");case "\u0073\u0073\u002e0\u0030":_dbd =_aaf .Round (_d .Second /100).AppendFormat (_dbd ,"\u0030\u0035\u002e0\u0030");case "\u0073\u0073\u002e\u0030\u0030\u0030":_dbd =_aaf .Round (_d .Second /1000).AppendFormat (_dbd ,"\u0030\u0035\u002e\u0030\u0030\u0030");case "\u0041\u004d\u002fP\u004d":_dbd =_aaf .AppendFormat (_dbd ,"\u0050\u004d");case "\u005b\u0068\u005d":_dbd =_df .AppendInt (_dbd ,int64 (_acff *24),10);case "\u005b\u006d\u005d":_dbd =_df .AppendInt (_dbd ,int64 (_acff *24*60),10);case "\u005b\u0073\u005d":_dbd =_df .AppendInt (_dbd ,int64 (_acff *24*60*60),10);case "":default:_ee .Log ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0069\u006d\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_fag );};if _cg [_agc ]==':'{_dbd =append (_dbd ,':');};};return _dbd ;};const _bfb int =34;

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String (v string ,f string )string {_gcb :=Parse (f );var _cd Format ;if len (_gcb )==1{_cd =_gcb [0];}else if len (_gcb )==4{_cd =_gcb [3];};_ggb :=false ;for _ ,_aaa :=range _cd .Whole {if _aaa .Type ==FmtTypeText {_ggb =true ;};};if !_ggb {return v ;};_dd :=_e .Buffer {};for _ ,_cea :=range _cd .Whole {switch _cea .Type {case FmtTypeLiteral :_dd .WriteByte (_cea .Literal );case FmtTypeText :_dd .WriteString (v );};};return _dd .String ();};const _gdb int =0;func (_acd *Lexer )nextFmt (){_acd ._cedb =append (_acd ._cedb ,_acd ._bada );_acd ._bada =Format {}};func _gf (_fc []byte )[]byte {for _aae :=0;_aae < len (_fc )/2;_aae ++{_ddg :=len (_fc )-1-_aae ;_fc [_aae ],_fc [_ddg ]=_fc [_ddg ],_fc [_aae ];};return _fc ;};const _cf ="\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0074\u0065\u0072a\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0069\u0067\u0069\u0074\u0046\u006d\u0074\u0054y\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0043o\u006d\u006d\u0061\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0065\u0063\u0069\u006da\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065Pe\u0072\u0063e\u006e\u0074\u0046\u006d\u0074\u0054\u0079\u0070e\u0044\u006f\u006c\u006c\u0061\u0072\u0046\u006d\u0074Ty\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0054\u0068\u006f\u0075\u0073\u0061n\u0064\u0073\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0055n\u0064\u0065\u0072\u0073c\u006f\u0072\u0065\u0046\u006d\u0074T\u0079\u0070\u0065\u0044\u0061\u0074\u0065\u0046\u006d\u0074\u0054y\u0070e\u0054\u0069\u006d\u0065\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0046\u0072\u0061\u0063t\u0069\u006f\u006e\u0046\u006dt\u0054\u0079\u0070\u0065\u0054e\u0078\u0074";

// Token is a format token in the Excel format string.
type Token struct{Type FmtType ;Literal byte ;DateTime string ;};const _ac =1e-10;func _gbe (_eef int64 )int64 {if _eef < 0{return -_eef ;};return _eef ;};func _af (_gcf ,_cc float64 ,_fd Format )[]byte {if len (_fd .Fractional )==0{return nil ;};_cbb :=_df .AppendFloat (nil ,_gcf ,'f',-1,64);if len (_cbb )> 2{_cbb =_cbb [2:];}else {_cbb =nil ;};_be :=make ([]byte ,0,len (_cbb ));_be =append (_be ,'.');_adgf :=0;_cca :for _gdf :=0;_gdf < len (_fd .Fractional );_gdf ++{_dg :=_gdf ;_dec :=_fd .Fractional [_gdf ];switch _dec .Type {case FmtTypeDigit :if _dg < len (_cbb ){_be =append (_be ,_cbb [_dg ]);_adgf ++;}else {_be =append (_be ,'0');};case FmtTypeDigitOpt :if _dg >=0{_be =append (_be ,_cbb [_dg ]);_adgf ++;}else {break _cca ;};case FmtTypeLiteral :_be =append (_be ,_dec .Literal );default:_ee .Log ("\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020f\u0072\u0061\u0063\u0074\u0069\u006f\u006ea\u006c\u0020\u0025\u0076",_dec );};};return _be ;};func _cbc (_badb int64 ,_eba Format )[]byte {if !_eba .IsExponential ||len (_eba .Exponent )==0{return nil ;};_gce :=_df .AppendInt (nil ,_gbe (_badb ),10);_bbc :=make ([]byte ,0,len (_gce )+2);_bbc =append (_bbc ,'E');if _badb >=0{_bbc =append (_bbc ,'+');}else {_bbc =append (_bbc ,'-');_badb *=-1;};_adgb :=0;_fda :for _ged :=len (_eba .Exponent )-1;_ged >=0;_ged --{_bbgg :=len (_gce )-1-_adgb ;_gea :=_eba .Exponent [_ged ];switch _gea .Type {case FmtTypeDigit :if _bbgg >=0{_bbc =append (_bbc ,_gce [_bbgg ]);_adgb ++;}else {_bbc =append (_bbc ,'0');};case FmtTypeDigitOpt :if _bbgg >=0{_bbc =append (_bbc ,_gce [_bbgg ]);_adgb ++;}else {for _abg :=_ged ;_abg >=0;_abg --{_ggc :=_eba .Exponent [_abg ];if _ggc .Type ==FmtTypeLiteral {_bbc =append (_bbc ,_ggc .Literal );};};break _fda ;};case FmtTypeLiteral :_bbc =append (_bbc ,_gea .Literal );default:_ee .Log ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0065\u0078p\u0020\u0025\u0076",_gea );};};if _adgb < len (_gce ){_bbc =append (_bbc ,_gce [len (_gce )-_adgb -1:_adgb -1]...);};_gf (_bbc [2:]);return _bbc ;};

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number (v float64 ,f string )string {if f ==""||f =="\u0047e\u006e\u0065\u0072\u0061\u006c"||f =="\u0040"{return NumberGeneric (v );};_fgd :=Parse (f );if len (_fgd )==1{return _bb (v ,_fgd [0],false );}else if len (_fgd )> 1&&v < 0{return _bb (v ,_fgd [1],true );}else if len (_fgd )> 2&&v ==0{return _bb (v ,_fgd [2],false );};return _bb (v ,_fgd [0],false );};var _gg =[...]uint8 {0,14,26,41,53,67,81,94,118,135,146,157,172,183};func _bad (_gbdd ,_eb float64 ,_bbg Format )[]byte {if len (_bbg .Whole )==0{return nil ;};_gbc :=_d .Date (1899,12,30,0,0,0,0,_d .UTC );_gbg :=_gbc .Add (_d .Duration (_eb *float64 (24*_d .Hour )));_gbg =_afd (_gbg );_dab :=_df .AppendFloat (nil ,_gbdd ,'f',-1,64);_ea :=make ([]byte ,0,len (_dab ));_bbgd :=0;_ced :=1;_gba :for _badf :=len (_bbg .Whole )-1;_badf >=0;_badf --{_bbgf :=len (_dab )-1-_bbgd ;_deg :=_bbg .Whole [_badf ];switch _deg .Type {case FmtTypeDigit :if _bbgf >=0{_ea =append (_ea ,_dab [_bbgf ]);_bbgd ++;_ced =_badf ;}else {_ea =append (_ea ,'0');};case FmtTypeDigitOpt :if _bbgf >=0{_ea =append (_ea ,_dab [_bbgf ]);_bbgd ++;_ced =_badf ;}else {for _fb :=_badf ;_fb >=0;_fb --{_fcg :=_bbg .Whole [_fb ];if _fcg .Type ==FmtTypeLiteral {_ea =append (_ea ,_fcg .Literal );};};break _gba ;};case FmtTypeDollar :for _eed :=_bbgd ;_eed < len (_dab );_eed ++{_ea =append (_ea ,_dab [len (_dab )-1-_eed ]);_bbgd ++;};_ea =append (_ea ,'$');case FmtTypeComma :if !_bbg ._db {_ea =append (_ea ,',');};case FmtTypeLiteral :_ea =append (_ea ,_deg .Literal );case FmtTypeDate :_ea =append (_ea ,_gf (_bafa (_gbg ,_deg .DateTime ))...);case FmtTypeTime :_ea =append (_ea ,_gf (_cce (_gbg ,_eb ,_deg .DateTime ))...);default:_ee .Log ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0077\u0068\u006f\u006c\u0065 \u0025\u0076",_deg );};};_ga :=_gf (_ea );if _bbgd < len (_dab )&&(_bbgd !=0||_bbg ._bd ){_ebf :=len (_dab )-_bbgd ;_bde :=make ([]byte ,len (_ga )+_ebf );copy (_bde ,_ga [0:_ced ]);copy (_bde [_ced :],_dab [0:]);copy (_bde [_ced +_ebf :],_ga [_ced :]);_ga =_bde ;};if _bbg ._db {_cfg :=_e .Buffer {};_ecg :=0;for _gfg :=len (_ga )-1;_gfg >=0;_gfg --{if !(_ga [_gfg ]>='0'&&_ga [_gfg ]<='9'){_ecg ++;}else {break ;};};for _fa :=0;_fa < len (_ga );_fa ++{_acg :=(len (_ga )-_fa -_ecg );if _acg %3==0&&_acg !=0&&_fa !=0{_cfg .WriteByte (',');};_cfg .WriteByte (_ga [_fa ]);};_ga =_cfg .Bytes ();};return _ga ;};

// Format is a parsed number format.
type Format struct{Whole []Token ;Fractional []Token ;Exponent []Token ;IsExponential bool ;_fg bool ;_gc bool ;_de bool ;_db bool ;_dc bool ;_bd bool ;_ce int64 ;_gb int ;};

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value (v string ,f string )string {if IsNumber (v ){_fe ,_ :=_df .ParseFloat (v ,64);return Number (_fe ,f );};return String (v ,f );};const _gbee int =0;type Lexer struct{_bada Format ;_cedb []Format ;};func _afd (_gacd _d .Time )_d .Time {_gacd =_gacd .UTC ();return _d .Date (_gacd .Year (),_gacd .Month (),_gacd .Day (),_gacd .Hour (),_gacd .Minute (),_gacd .Second (),_gacd .Nanosecond (),_d .Local );};const _aab int =-1;func _baf (_gbf []byte )[]byte {for _eff :=len (_gbf )-1;_eff > 0;_eff --{if _gbf [_eff ]=='9'+1{_gbf [_eff ]='0';if _gbf [_eff -1]=='.'{_eff --;};_gbf [_eff -1]++;};};if _gbf [0]=='9'+1{_gbf [0]='0';copy (_gbf [1:],_gbf [0:]);_gbf [0]='1';};return _gbf ;};func _egf (_fae []byte )[]byte {_agd :=len (_fae );_eea :=false ;_dabe :=false ;for _dbb :=len (_fae )-1;_dbb >=0;_dbb --{if _fae [_dbb ]=='0'&&!_dabe &&!_eea {_agd =_dbb ;}else if _fae [_dbb ]=='.'{_eea =true ;}else {_dabe =true ;};};if _eea &&_dabe {if _fae [_agd -1]=='.'{_agd --;};return _fae [0:_agd ];};return _fae ;};func _bafa (_fbf _d .Time ,_eaa string )[]byte {_agf :=[]byte {};_efb :=0;for _cbg :=0;_cbg < len (_eaa );_cbg ++{var _gdde string ;if _eaa [_cbg ]=='/'{_gdde =string (_eaa [_efb :_cbg ]);_efb =_cbg +1;}else if _cbg ==len (_eaa )-1{_gdde =string (_eaa [_efb :_cbg +1]);}else {continue ;};switch _gdde {case "\u0079\u0079":_agf =_fbf .AppendFormat (_agf ,"\u0030\u0036");case "\u0079\u0079\u0079\u0079":_agf =_fbf .AppendFormat (_agf ,"\u0032\u0030\u0030\u0036");case "\u006d":_agf =_fbf .AppendFormat (_agf ,"\u0031");case "\u006d\u006d":_agf =_fbf .AppendFormat (_agf ,"\u0030\u0031");case "\u006d\u006d\u006d":_agf =_fbf .AppendFormat (_agf ,"\u004a\u0061\u006e");case "\u006d\u006d\u006d\u006d":_agf =_fbf .AppendFormat (_agf ,"\u004aa\u006e\u0075\u0061\u0072\u0079");case "\u006d\u006d\u006dm\u006d":switch _fbf .Month (){case _d .January ,_d .July ,_d .June :_agf =append (_agf ,'J');case _d .February :_agf =append (_agf ,'M');case _d .March ,_d .May :_agf =append (_agf ,'M');case _d .April ,_d .August :_agf =append (_agf ,'A');case _d .September :_agf =append (_agf ,'S');case _d .October :_agf =append (_agf ,'O');case _d .November :_agf =append (_agf ,'N');case _d .December :_agf =append (_agf ,'D');};case "\u0064":_agf =_fbf .AppendFormat (_agf ,"\u0032");case "\u0064\u0064":_agf =_fbf .AppendFormat (_agf ,"\u0030\u0032");case "\u0064\u0064\u0064":_agf =_fbf .AppendFormat (_agf ,"\u004d\u006f\u006e");case "\u0064\u0064\u0064\u0064":_agf =_fbf .AppendFormat (_agf ,"\u004d\u006f\u006e\u0064\u0061\u0079");default:_ee .Log ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_gdde );};if _eaa [_cbg ]=='/'{_agf =append (_agf ,'/');};};return _agf ;};const _eaf int =-1;func (_bff *Lexer )Lex (r _c .Reader ){_cga ,_fgbd ,_cfga :=0,0,0;_feb :=-1;_gge ,_eaaa ,_efe :=0,0,0;_ =_eaaa ;_ =_efe ;_bdf :=1;_ =_bdf ;_cdd :=make ([]byte ,4096);_ed :=false ;for !_ed {_edf :=0;if _gge > 0{_edf =_fgbd -_gge ;};_fgbd =0;_cdg ,_gfge :=r .Read (_cdd [_edf :]);if _cdg ==0||_gfge !=nil {_ed =true ;};_cfga =_cdg +_edf ;if _cfga < len (_cdd ){_feb =_cfga ;};{_cga =_dbf ;_gge =0;_eaaa =0;_efe =0;};{if _fgbd ==_cfga {goto _fede ;};switch _cga {case 34:goto _cdb ;case 35:goto _bcf ;case 0:goto _ffbb ;case 36:goto _dcd ;case 37:goto _fbe ;case 1:goto _cecd ;case 2:goto _ggd ;case 38:goto _dfc ;case 3:goto _defb ;case 4:goto _abe ;case 39:goto _gag ;case 5:goto _ddgb ;case 6:goto _acag ;case 7:goto _cbd ;case 8:goto _eeg ;case 40:goto _egc ;case 9:goto _fgdb ;case 41:goto _ebbf ;case 10:goto _bbf ;case 42:goto _gac ;case 11:goto _ecge ;case 43:goto _afc ;case 44:goto _ccd ;case 45:goto _dbfg ;case 12:goto _fcd ;case 46:goto _dbbb ;case 13:goto _eefc ;case 14:goto _gbga ;case 15:goto _adba ;case 16:goto _bea ;case 47:goto _dae ;case 17:goto _afb ;case 48:goto _ggce ;case 18:goto _bdgd ;case 19:goto _aea ;case 20:goto _aeg ;case 49:goto _bcb ;case 50:goto _fcfe ;case 21:goto _bed ;case 22:goto _abdb ;case 23:goto _gbcc ;case 24:goto _beb ;case 25:goto _gcff ;case 51:goto _efg ;case 26:goto _fceb ;case 52:goto _dgba ;case 53:goto _fecb ;case 54:goto _cfd ;case 55:goto _dfd ;case 56:goto _agab ;case 57:goto _eag ;case 27:goto _gde ;case 28:goto _eccg ;case 29:goto _gfef ;case 30:goto _dbbba ;case 31:goto _acfff ;case 58:goto _cagf ;case 32:goto _bfd ;case 59:goto _cfaed ;case 33:goto _fcca ;case 60:goto _edg ;case 61:goto _fbeb ;case 62:goto _gfbe ;};goto _cfaeb ;_bcc :switch _efe {case 2:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeDigit ,nil );};case 3:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeDigitOpt ,nil );};case 5:{_fgbd =(_eaaa )-1;};case 8:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypePercent ,nil );};case 13:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeFraction ,_cdd [_gge :_eaaa ]);};case 14:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeDate ,_cdd [_gge :_eaaa ]);};case 15:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeTime ,_cdd [_gge :_eaaa ]);};case 16:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeTime ,_cdd [_gge :_eaaa ]);};case 18:{_fgbd =(_eaaa )-1;};case 20:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge :_eaaa ]);};case 21:{_fgbd =(_eaaa )-1;_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge +1:_eaaa -1]);};};goto _gedf ;_ecgb :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeFraction ,_cdd [_gge :_eaaa ]);};goto _gedf ;_adge :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeDigitOpt ,nil );};goto _gedf ;_bdc :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeDigitOptThousands ,nil );};goto _gedf ;_bffc :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypePercent ,nil );};goto _gedf ;_agfa :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeDate ,_cdd [_gge :_eaaa ]);};goto _gedf ;_bac :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeDigit ,nil );};goto _gedf ;_dcf :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeTime ,_cdd [_gge :_eaaa ]);};goto _gedf ;_fad :_fgbd =(_eaaa )-1;{_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge :_eaaa ]);};goto _gedf ;_fcc :_eaaa =_fgbd +1;{_bff ._bada ._de =true ;};goto _gedf ;_fcga :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge :_eaaa ]);};goto _gedf ;_abgb :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeDollar ,nil );};goto _gedf ;_gddf :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeComma ,nil );};goto _gedf ;_agg :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeDecimal ,nil );};goto _gedf ;_acdc :_eaaa =_fgbd +1;{_bff .nextFmt ();};goto _gedf ;_efa :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeText ,nil );};goto _gedf ;_gee :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeUnderscore ,nil );};goto _gedf ;_acb :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge :_eaaa ]);};goto _gedf ;_bbgdb :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge +1:_eaaa -1]);};goto _gedf ;_dfg :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeDigitOpt ,nil );};goto _gedf ;_fec :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeFraction ,_cdd [_gge :_eaaa ]);};goto _gedf ;_bge :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypePercent ,nil );};goto _gedf ;_bce :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeDate ,_cdd [_gge :_eaaa ]);};goto _gedf ;_add :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeDigit ,nil );};goto _gedf ;_fgf :_eaaa =_fgbd ;_fgbd --;{_bff ._bada .AddToken (FmtTypeTime ,_cdd [_gge :_eaaa ]);};goto _gedf ;_ccc :_eaaa =_fgbd ;_fgbd --;{};goto _gedf ;_bae :_eaaa =_fgbd +1;{_bff ._bada .IsExponential =true ;};goto _gedf ;_dbdc :_eaaa =_fgbd +1;{_bff ._bada .AddToken (FmtTypeLiteral ,_cdd [_gge +1:_eaaa ]);};goto _gedf ;_gedf :_gge =0;if _fgbd ++;_fgbd ==_cfga {goto _egea ;};_cdb :_gge =_fgbd ;switch _cdd [_fgbd ]{case 34:goto _ceg ;case 35:goto _fdb ;case 36:goto _abgb ;case 37:goto _cab ;case 44:goto _gddf ;case 46:goto _agg ;case 47:goto _aebb ;case 48:goto _afg ;case 58:goto _dgbe ;case 59:goto _acdc ;case 63:goto _bcg ;case 64:goto _efa ;case 65:goto _bcag ;case 69:goto _baca ;case 71:goto _gdc ;case 91:goto _faef ;case 92:goto _ded ;case 95:goto _gee ;case 100:goto _aebb ;case 104:goto _dgbe ;case 109:goto _gded ;case 115:goto _egcb ;case 121:goto _bdcc ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbg ;};goto _fcga ;_ceg :_eaaa =_fgbd +1;_efe =20;goto _gead ;_gead :if _fgbd ++;_fgbd ==_cfga {goto _fdf ;};_bcf :if _cdd [_fgbd ]==34{goto _bda ;};goto _bafc ;_bafc :if _fgbd ++;_fgbd ==_cfga {goto _bfg ;};_ffbb :if _cdd [_fgbd ]==34{goto _bda ;};goto _bafc ;_bda :_eaaa =_fgbd +1;_efe =21;goto _gad ;_gad :if _fgbd ++;_fgbd ==_cfga {goto _gacg ;};_dcd :if _cdd [_fgbd ]==34{goto _bafc ;};goto _bbgdb ;_fdb :_eaaa =_fgbd +1;_efe =3;goto _ecc ;_ecc :if _fgbd ++;_fgbd ==_cfga {goto _cef ;};_fbe :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cfae ;case 44:goto _fef ;case 47:goto _fcf ;case 48:goto _cfae ;case 63:goto _cfae ;};goto _dfg ;_cfae :if _fgbd ++;_fgbd ==_cfga {goto _gcbg ;};_cecd :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cfae ;case 47:goto _fcf ;case 48:goto _cfae ;case 63:goto _cfae ;};goto _bcc ;_fcf :if _fgbd ++;_fgbd ==_cfga {goto _dde ;};_ggd :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _ccb ;case 48:goto _dce ;case 63:goto _gbdb ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _gfd ;};goto _bcc ;_gbdb :_eaaa =_fgbd +1;goto _ebbe ;_ebbe :if _fgbd ++;_fgbd ==_cfga {goto _acc ;};_dfc :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _gbdb ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _gbdb ;case 63:goto _gbdb ;case 65:goto _aga ;};goto _fec ;_aga :if _fgbd ++;_fgbd ==_cfga {goto _gbgf ;};_defb :switch _cdd [_fgbd ]{case 47:goto _fgba ;case 77:goto _ebd ;};goto _ecgb ;_fgba :if _fgbd ++;_fgbd ==_cfga {goto _bafac ;};_abe :if _cdd [_fgbd ]==80{goto _acdf ;};goto _ecgb ;_acdf :_eaaa =_fgbd +1;goto _abdd ;_abdd :if _fgbd ++;_fgbd ==_cfga {goto _cge ;};_gag :if _cdd [_fgbd ]==65{goto _aga ;};goto _fec ;_ebd :if _fgbd ++;_fgbd ==_cfga {goto _fab ;};_ddgb :if _cdd [_fgbd ]==47{goto _fegc ;};goto _ecgb ;_fegc :if _fgbd ++;_fgbd ==_cfga {goto _bbd ;};_acag :if _cdd [_fgbd ]==80{goto _ca ;};goto _ecgb ;_ca :if _fgbd ++;_fgbd ==_cfga {goto _ebbb ;};_cbd :if _cdd [_fgbd ]==77{goto _acdf ;};goto _ecgb ;_ccb :if _fgbd ++;_fgbd ==_cfga {goto _agce ;};_eeg :switch _cdd [_fgbd ]{case 35:goto _faec ;case 37:goto _cgf ;case 63:goto _faec ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbc ;};goto _bcc ;_faec :_eaaa =_fgbd +1;goto _ege ;_ege :if _fgbd ++;_fgbd ==_cfga {goto _ggef ;};_egc :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _fed ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _gbdb ;case 63:goto _gbdb ;case 65:goto _aga ;};goto _fec ;_fed :if _fgbd ++;_fgbd ==_cfga {goto _ddb ;};_fgdb :switch _cdd [_fgbd ]{case 35:goto _bffcd ;case 44:goto _bffcd ;case 46:goto _bffcd ;case 48:goto _bffcd ;case 63:goto _bffcd ;};goto _ecgb ;_bffcd :_eaaa =_fgbd +1;goto _fcgd ;_fcgd :if _fgbd ++;_fgbd ==_cfga {goto _bdcg ;};_ebbf :switch _cdd [_fgbd ]{case 35:goto _bffcd ;case 44:goto _bffcd ;case 46:goto _bffcd ;case 48:goto _bffcd ;case 63:goto _bffcd ;case 65:goto _aga ;};goto _fec ;_cgf :if _fgbd ++;_fgbd ==_cfga {goto _dfdc ;};_bbf :if _cdd [_fgbd ]==37{goto _cgf ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbc ;};goto _bcc ;_fbc :_eaaa =_fgbd +1;_efe =13;goto _ceb ;_ceb :if _fgbd ++;_fgbd ==_cfga {goto _afbg ;};_gac :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _dgb ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _aaaa ;case 63:goto _gbdb ;case 65:goto _aga ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbc ;};goto _fec ;_dgb :if _fgbd ++;_fgbd ==_cfga {goto _faa ;};_ecge :switch _cdd [_fgbd ]{case 35:goto _bffcd ;case 37:goto _cgf ;case 44:goto _bffcd ;case 46:goto _bffcd ;case 63:goto _bffcd ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbc ;};goto _ecgb ;_aaaa :_eaaa =_fgbd +1;goto _edb ;_edb :if _fgbd ++;_fgbd ==_cfga {goto _bced ;};_afc :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _aaaa ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _aaaa ;case 63:goto _gbdb ;case 65:goto _aga ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fbc ;};goto _fec ;_dce :_eaaa =_fgbd +1;goto _bffcc ;_bffcc :if _fgbd ++;_fgbd ==_cfga {goto _bffa ;};_ccd :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _aaaa ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _dce ;case 63:goto _gbdb ;case 65:goto _aga ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _gfd ;};goto _fec ;_gfd :_eaaa =_fgbd +1;goto _bag ;_bag :if _fgbd ++;_fgbd ==_cfga {goto _aggf ;};_dbfg :switch _cdd [_fgbd ]{case 35:goto _gbdb ;case 37:goto _fbc ;case 44:goto _gbdb ;case 46:goto _gbdb ;case 48:goto _dce ;case 63:goto _gbdb ;case 65:goto _aga ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _gfd ;};goto _fec ;_fef :if _fgbd ++;_fgbd ==_cfga {goto _cff ;};_fcd :if _cdd [_fgbd ]==35{goto _bdc ;};goto _adge ;_cab :_eaaa =_fgbd +1;_efe =8;goto _cag ;_cag :if _fgbd ++;_fgbd ==_cfga {goto _gbdbb ;};_dbbb :switch _cdd [_fgbd ]{case 35:goto _fddf ;case 37:goto _fcfb ;case 48:goto _eca ;case 63:goto _fddf ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fccb ;};goto _bge ;_fddf :if _fgbd ++;_fgbd ==_cfga {goto _fac ;};_eefc :switch _cdd [_fgbd ]{case 35:goto _fddf ;case 47:goto _fcf ;case 48:goto _fddf ;case 63:goto _fddf ;};goto _bffc ;_fcfb :if _fgbd ++;_fgbd ==_cfga {goto _fcb ;};_gbga :if _cdd [_fgbd ]==37{goto _fcfb ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fccb ;};goto _bcc ;_fccb :if _fgbd ++;_fgbd ==_cfga {goto _gbca ;};_adba :switch _cdd [_fgbd ]{case 37:goto _fcfb ;case 47:goto _fcf ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fccb ;};goto _bcc ;_eca :if _fgbd ++;_fgbd ==_cfga {goto _ddc ;};_bea :switch _cdd [_fgbd ]{case 35:goto _fddf ;case 37:goto _fcfb ;case 47:goto _fcf ;case 48:goto _eca ;case 63:goto _fddf ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fccb ;};goto _bffc ;_aebb :_eaaa =_fgbd +1;goto _aff ;_aff :if _fgbd ++;_fgbd ==_cfga {goto _egbf ;};_dae :switch _cdd [_fgbd ]{case 47:goto _aebb ;case 100:goto _aebb ;case 109:goto _aebb ;case 121:goto _bbb ;};goto _bce ;_bbb :if _fgbd ++;_fgbd ==_cfga {goto _ecf ;};_afb :if _cdd [_fgbd ]==121{goto _aebb ;};goto _agfa ;_afg :_eaaa =_fgbd +1;_efe =2;goto _ega ;_ega :if _fgbd ++;_fgbd ==_cfga {goto _gec ;};_ggce :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cgd ;case 47:goto _fcf ;case 48:goto _cddb ;case 63:goto _cfae ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _cede ;};goto _add ;_cgd :if _fgbd ++;_fgbd ==_cfga {goto _egdf ;};_bdgd :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cgd ;case 47:goto _fcf ;case 48:goto _cgd ;case 63:goto _cfae ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _fccb ;};goto _bac ;_cddb :if _fgbd ++;_fgbd ==_cfga {goto _effb ;};_aea :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cgd ;case 47:goto _fcf ;case 48:goto _cddb ;case 63:goto _cfae ;};if 49<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _cede ;};goto _bac ;_cede :if _fgbd ++;_fgbd ==_cfga {goto _dba ;};_aeg :switch _cdd [_fgbd ]{case 37:goto _fccb ;case 47:goto _fcf ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _cede ;};goto _bcc ;_fbg :_eaaa =_fgbd +1;_efe =20;goto _badfb ;_badfb :if _fgbd ++;_fgbd ==_cfga {goto _dcfg ;};_bcb :switch _cdd [_fgbd ]{case 37:goto _fccb ;case 47:goto _fcf ;};if 48<=_cdd [_fgbd ]&&_cdd [_fgbd ]<=57{goto _cede ;};goto _acb ;_dgbe :_eaaa =_fgbd +1;_efe =15;goto _ggdg ;_ggdg :if _fgbd ++;_fgbd ==_cfga {goto _ggff ;};_fcfe :switch _cdd [_fgbd ]{case 58:goto _dgbe ;case 65:goto _cdgd ;case 104:goto _dgbe ;case 109:goto _dgbe ;case 115:goto _egcb ;};goto _fgf ;_cdgd :if _fgbd ++;_fgbd ==_cfga {goto _fegd ;};_bed :switch _cdd [_fgbd ]{case 47:goto _acaf ;case 77:goto _bca ;};goto _bcc ;_acaf :if _fgbd ++;_fgbd ==_cfga {goto _gged ;};_abdb :if _cdd [_fgbd ]==80{goto _dgbe ;};goto _bcc ;_bca :if _fgbd ++;_fgbd ==_cfga {goto _degf ;};_gbcc :if _cdd [_fgbd ]==47{goto _dad ;};goto _bcc ;_dad :if _fgbd ++;_fgbd ==_cfga {goto _cdc ;};_beb :if _cdd [_fgbd ]==80{goto _effa ;};goto _bcc ;_effa :if _fgbd ++;_fgbd ==_cfga {goto _fagg ;};_gcff :if _cdd [_fgbd ]==77{goto _dgbe ;};goto _bcc ;_egcb :_eaaa =_fgbd +1;_efe =15;goto _egb ;_egb :if _fgbd ++;_fgbd ==_cfga {goto _feeb ;};_efg :switch _cdd [_fgbd ]{case 46:goto _ffc ;case 58:goto _dgbe ;case 65:goto _cdgd ;case 104:goto _dgbe ;case 109:goto _dgbe ;case 115:goto _egcb ;};goto _fgf ;_ffc :if _fgbd ++;_fgbd ==_cfga {goto _faf ;};_fceb :if _cdd [_fgbd ]==48{goto _edbc ;};goto _dcf ;_edbc :_eaaa =_fgbd +1;_efe =15;goto _efae ;_efae :if _fgbd ++;_fgbd ==_cfga {goto _geee ;};_dgba :switch _cdd [_fgbd ]{case 48:goto _ffcd ;case 58:goto _dgbe ;case 65:goto _cdgd ;case 104:goto _dgbe ;case 109:goto _dgbe ;case 115:goto _egcb ;};goto _fgf ;_ffcd :_eaaa =_fgbd +1;_efe =15;goto _fbca ;_fbca :if _fgbd ++;_fgbd ==_cfga {goto _accf ;};_fecb :switch _cdd [_fgbd ]{case 48:goto _dgbe ;case 58:goto _dgbe ;case 65:goto _cdgd ;case 104:goto _dgbe ;case 109:goto _dgbe ;case 115:goto _egcb ;};goto _fgf ;_bcg :_eaaa =_fgbd +1;_efe =5;goto _cegg ;_cegg :if _fgbd ++;_fgbd ==_cfga {goto _bbdg ;};_cfd :switch _cdd [_fgbd ]{case 35:goto _cfae ;case 37:goto _cfae ;case 47:goto _fcf ;case 48:goto _cfae ;case 63:goto _cfae ;};goto _ccc ;_bcag :_eaaa =_fgbd +1;_efe =20;goto _bfe ;_bfe :if _fgbd ++;_fgbd ==_cfga {goto _eec ;};_dfd :switch _cdd [_fgbd ]{case 47:goto _acaf ;case 77:goto _bca ;};goto _acb ;_baca :if _fgbd ++;_fgbd ==_cfga {goto _ebbc ;};_agab :switch _cdd [_fgbd ]{case 43:goto _bae ;case 45:goto _bae ;};goto _acb ;_gdc :_eaaa =_fgbd +1;goto _aefe ;_aefe :if _fgbd ++;_fgbd ==_cfga {goto _gddd ;};_eag :if _cdd [_fgbd ]==101{goto _bgf ;};goto _acb ;_bgf :if _fgbd ++;_fgbd ==_cfga {goto _efc ;};_gde :if _cdd [_fgbd ]==110{goto _egaa ;};goto _fad ;_egaa :if _fgbd ++;_fgbd ==_cfga {goto _baed ;};_eccg :if _cdd [_fgbd ]==101{goto _cdf ;};goto _fad ;_cdf :if _fgbd ++;_fgbd ==_cfga {goto _gbag ;};_gfef :if _cdd [_fgbd ]==114{goto _bfff ;};goto _fad ;_bfff :if _fgbd ++;_fgbd ==_cfga {goto _cgaa ;};_dbbba :if _cdd [_fgbd ]==97{goto _cfb ;};goto _fad ;_cfb :if _fgbd ++;_fgbd ==_cfga {goto _gcba ;};_acfff :if _cdd [_fgbd ]==108{goto _fcc ;};goto _fad ;_faef :_eaaa =_fgbd +1;_efe =20;goto _badd ;_badd :if _fgbd ++;_fgbd ==_cfga {goto _faab ;};_cagf :switch _cdd [_fgbd ]{case 104:goto _cgff ;case 109:goto _cgff ;case 115:goto _cgff ;};goto _cgfa ;_cgfa :if _fgbd ++;_fgbd ==_cfga {goto _dgd ;};_bfd :if _cdd [_fgbd ]==93{goto _bacf ;};goto _cgfa ;_bacf :_eaaa =_fgbd +1;_efe =18;goto _faecd ;_ccdb :_eaaa =_fgbd +1;_efe =16;goto _faecd ;_faecd :if _fgbd ++;_fgbd ==_cfga {goto _bccc ;};_cfaed :if _cdd [_fgbd ]==93{goto _bacf ;};goto _cgfa ;_cgff :if _fgbd ++;_fgbd ==_cfga {goto _bcccd ;};_fcca :if _cdd [_fgbd ]==93{goto _ccdb ;};goto _cgfa ;_ded :if _fgbd ++;_fgbd ==_cfga {goto _ade ;};_edg :goto _dbdc ;_gded :_eaaa =_fgbd +1;_efe =14;goto _ecca ;_ecca :if _fgbd ++;_fgbd ==_cfga {goto _cege ;};_fbeb :switch _cdd [_fgbd ]{case 47:goto _aebb ;case 58:goto _dgbe ;case 65:goto _cdgd ;case 100:goto _aebb ;case 104:goto _dgbe ;case 109:goto _gded ;case 115:goto _egcb ;case 121:goto _bbb ;};goto _bce ;_bdcc :if _fgbd ++;_fgbd ==_cfga {goto _egcg ;};_gfbe :if _cdd [_fgbd ]==121{goto _aebb ;};goto _acb ;_cfaeb :_egea :_cga =34;goto _fede ;_fdf :_cga =35;goto _fede ;_bfg :_cga =0;goto _fede ;_gacg :_cga =36;goto _fede ;_cef :_cga =37;goto _fede ;_gcbg :_cga =1;goto _fede ;_dde :_cga =2;goto _fede ;_acc :_cga =38;goto _fede ;_gbgf :_cga =3;goto _fede ;_bafac :_cga =4;goto _fede ;_cge :_cga =39;goto _fede ;_fab :_cga =5;goto _fede ;_bbd :_cga =6;goto _fede ;_ebbb :_cga =7;goto _fede ;_agce :_cga =8;goto _fede ;_ggef :_cga =40;goto _fede ;_ddb :_cga =9;goto _fede ;_bdcg :_cga =41;goto _fede ;_dfdc :_cga =10;goto _fede ;_afbg :_cga =42;goto _fede ;_faa :_cga =11;goto _fede ;_bced :_cga =43;goto _fede ;_bffa :_cga =44;goto _fede ;_aggf :_cga =45;goto _fede ;_cff :_cga =12;goto _fede ;_gbdbb :_cga =46;goto _fede ;_fac :_cga =13;goto _fede ;_fcb :_cga =14;goto _fede ;_gbca :_cga =15;goto _fede ;_ddc :_cga =16;goto _fede ;_egbf :_cga =47;goto _fede ;_ecf :_cga =17;goto _fede ;_gec :_cga =48;goto _fede ;_egdf :_cga =18;goto _fede ;_effb :_cga =19;goto _fede ;_dba :_cga =20;goto _fede ;_dcfg :_cga =49;goto _fede ;_ggff :_cga =50;goto _fede ;_fegd :_cga =21;goto _fede ;_gged :_cga =22;goto _fede ;_degf :_cga =23;goto _fede ;_cdc :_cga =24;goto _fede ;_fagg :_cga =25;goto _fede ;_feeb :_cga =51;goto _fede ;_faf :_cga =26;goto _fede ;_geee :_cga =52;goto _fede ;_accf :_cga =53;goto _fede ;_bbdg :_cga =54;goto _fede ;_eec :_cga =55;goto _fede ;_ebbc :_cga =56;goto _fede ;_gddd :_cga =57;goto _fede ;_efc :_cga =27;goto _fede ;_baed :_cga =28;goto _fede ;_gbag :_cga =29;goto _fede ;_cgaa :_cga =30;goto _fede ;_gcba :_cga =31;goto _fede ;_faab :_cga =58;goto _fede ;_dgd :_cga =32;goto _fede ;_bccc :_cga =59;goto _fede ;_bcccd :_cga =33;goto _fede ;_ade :_cga =60;goto _fede ;_cege :_cga =61;goto _fede ;_egcg :_cga =62;goto _fede ;_fede :{};if _fgbd ==_feb {switch _cga {case 35:goto _acb ;case 0:goto _bcc ;case 36:goto _bbgdb ;case 37:goto _dfg ;case 1:goto _bcc ;case 2:goto _bcc ;case 38:goto _fec ;case 3:goto _ecgb ;case 4:goto _ecgb ;case 39:goto _fec ;case 5:goto _ecgb ;case 6:goto _ecgb ;case 7:goto _ecgb ;case 8:goto _bcc ;case 40:goto _fec ;case 9:goto _ecgb ;case 41:goto _fec ;case 10:goto _bcc ;case 42:goto _fec ;case 11:goto _ecgb ;case 43:goto _fec ;case 44:goto _fec ;case 45:goto _fec ;case 12:goto _adge ;case 46:goto _bge ;case 13:goto _bffc ;case 14:goto _bcc ;case 15:goto _bcc ;case 16:goto _bffc ;case 47:goto _bce ;case 17:goto _agfa ;case 48:goto _add ;case 18:goto _bac ;case 19:goto _bac ;case 20:goto _bcc ;case 49:goto _acb ;case 50:goto _fgf ;case 21:goto _bcc ;case 22:goto _bcc ;case 23:goto _bcc ;case 24:goto _bcc ;case 25:goto _bcc ;case 51:goto _fgf ;case 26:goto _dcf ;case 52:goto _fgf ;case 53:goto _fgf ;case 54:goto _ccc ;case 55:goto _acb ;case 56:goto _acb ;case 57:goto _acb ;case 27:goto _fad ;case 28:goto _fad ;case 29:goto _fad ;case 30:goto _fad ;case 31:goto _fad ;case 58:goto _acb ;case 32:goto _bcc ;case 59:goto _bcc ;case 33:goto _fad ;case 60:goto _acb ;case 61:goto _bce ;case 62:goto _acb ;};};};if _gge > 0{copy (_cdd [0:],_cdd [_gge :]);};};_ =_feb ;if _cga ==_eaf {_ee .Log ("\u0066o\u0072m\u0061\u0074\u0020\u0070\u0061r\u0073\u0065 \u0065\u0072\u0072\u006f\u0072");};};

// FmtType is the type of a format token.
//go:generate stringer -type=FmtType
type FmtType byte ;

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric (v float64 )string {if _f .Abs (v )>=_cb ||_f .Abs (v )<=_ac &&v !=0{return _bgb (v );};_daa :=make ([]byte ,0,15);_daa =_df .AppendFloat (_daa ,v ,'f',-1,64);if len (_daa )> 11{_gaa :=_daa [11]-'0';if _gaa >=5&&_gaa <=9{_daa [10]++;_daa =_daa [0:11];_daa =_baf (_daa );};_daa =_daa [0:11];}else if len (_daa )==11{if _daa [len (_daa )-1]=='9'{_daa [len (_daa )-1]++;_daa =_baf (_daa );};};_daa =_egf (_daa );return string (_daa );};func _bb (_cfaf float64 ,_ad Format ,_gfe bool )string {if _ad ._de {return NumberGeneric (_cfaf );};_ef :=make ([]byte ,0,20);_gdd :=_f .Signbit (_cfaf );_feg :=_f .Abs (_cfaf );_bg :=int64 (0);_bbe :=int64 (0);if _ad .IsExponential {for _feg >=10{_bbe ++;_feg /=10;};for _feg < 1{_bbe --;_feg *=10;};}else if _ad ._gc {_feg *=100;}else if _ad ._fg {if _ad ._ce ==0{_adg :=_f .Pow (10,float64 (_ad ._gb ));_fgb ,_eg :=1.0,1.0;_ =_fgb ;for _gfb :=1.0;_gfb < _adg ;_gfb ++{_ ,_bdb :=_f .Modf (_feg *float64 (_gfb ));if _bdb < _eg {_eg =_bdb ;_fgb =_gfb ;if _bdb ==0{break ;};};};_ad ._ce =int64 (_fgb );};_bg =int64 (_feg *float64 (_ad ._ce )+0.5);if len (_ad .Whole )> 0&&_bg > _ad ._ce {_bg =int64 (_feg *float64 (_ad ._ce ))%_ad ._ce ;_feg -=float64 (_bg )/float64 (_ad ._ce );}else {_feg -=float64 (_bg )/float64 (_ad ._ce );if _f .Abs (_feg )< 1{_ge :=true ;for _ ,_ff :=range _ad .Whole {if _ff .Type ==FmtTypeDigitOpt {continue ;};if _ff .Type ==FmtTypeLiteral &&_ff .Literal ==' '{continue ;};_ge =false ;};if _ge {_ad .Whole =nil ;};};};};_abc :=1;for _ ,_fce :=range _ad .Fractional {if _fce .Type ==FmtTypeDigit ||_fce .Type ==FmtTypeDigitOpt {_abc ++;};};_feg +=5*_f .Pow10 (-_abc );_gbd ,_def :=_f .Modf (_feg );_ef =append (_ef ,_bad (_gbd ,_cfaf ,_ad )...);_ef =append (_ef ,_af (_def ,_cfaf ,_ad )...);_ef =append (_ef ,_cbc (_bbe ,_ad )...);if _ad ._fg {_ef =_df .AppendInt (_ef ,_bg ,10);_ef =append (_ef ,'/');_ef =_df .AppendInt (_ef ,_ad ._ce ,10);};if !_gfe &&_gdd {return "\u002d"+string (_ef );};return string (_ef );};const _bc int =0;const _dbf int =34;