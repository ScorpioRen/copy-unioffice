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

package drawing ;import (_a "github.com/ScorpioRen/copy-unioffice";_c "github.com/ScorpioRen/copy-unioffice/color";_da "github.com/ScorpioRen/copy-unioffice/measurement";_aa "github.com/ScorpioRen/copy-unioffice/schema/soo/dml";);func (_dfg ShapeProperties )ensureXfrm (){if _dfg ._fc .Xfrm ==nil {_dfg ._fc .Xfrm =_aa .NewCT_Transform2D ();};};func (_eb LineProperties )SetNoFill (){_eb .clearFill ();_eb ._g .NoFill =_aa .NewCT_NoFillProperties ()};

// SetNumbered controls if bullets are numbered or not.
func (_fe ParagraphProperties )SetNumbered (scheme _aa .ST_TextAutonumberScheme ){if scheme ==_aa .ST_TextAutonumberSchemeUnset {_fe ._dff .BuAutoNum =nil ;}else {_fe ._dff .BuAutoNum =_aa .NewCT_TextAutonumberBullet ();_fe ._dff .BuAutoNum .TypeAttr =scheme ;};};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_aa .EG_TextRun )Run {return Run {x }};

// SetBulletFont controls the font for the bullet character.
func (_ac ParagraphProperties )SetBulletFont (f string ){if f ==""{_ac ._dff .BuFont =nil ;}else {_ac ._dff .BuFont =_aa .NewCT_TextFont ();_ac ._dff .BuFont .TypefaceAttr =f ;};};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_aa .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};type ShapeProperties struct{_fc *_aa .CT_ShapeProperties };

// SetAlign controls the paragraph alignment
func (_aaa ParagraphProperties )SetAlign (a _aa .ST_TextAlignType ){_aaa ._dff .AlgnAttr =a };

// LineJoin is the type of line join
type LineJoin byte ;

// SetPosition sets the position of the shape.
func (_fb ShapeProperties )SetPosition (x ,y _da .Distance ){_fb .ensureXfrm ();if _fb ._fc .Xfrm .Off ==nil {_fb ._fc .Xfrm .Off =_aa .NewCT_Point2D ();};_fb ._fc .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_a .Int64 (int64 (x /_da .EMU ));_fb ._fc .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_a .Int64 (int64 (y /_da .EMU ));};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_dff *_aa .CT_TextParagraphProperties ;};func (_bc ShapeProperties )SetNoFill (){_bc .clearFill ();_bc ._fc .NoFill =_aa .NewCT_NoFillProperties ();};func (_fcb ShapeProperties )clearFill (){_fcb ._fc .NoFill =nil ;_fcb ._fc .BlipFill =nil ;_fcb ._fc .GradFill =nil ;_fcb ._fc .GrpFill =nil ;_fcb ._fc .SolidFill =nil ;_fcb ._fc .PattFill =nil ;};

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_bb ShapeProperties )SetFlipHorizontal (b bool ){_bb .ensureXfrm ();if !b {_bb ._fc .Xfrm .FlipHAttr =nil ;}else {_bb ._fc .Xfrm .FlipHAttr =_a .Bool (true );};};

// SetFont controls the font of a run.
func (_gaa RunProperties )SetFont (s string ){_gaa ._cge .Latin =_aa .NewCT_TextFont ();_gaa ._cge .Latin .TypefaceAttr =s ;};type LineProperties struct{_g *_aa .CT_LineProperties };

// AddBreak adds a new line break to a paragraph.
func (_b Paragraph )AddBreak (){_fa :=_aa .NewEG_TextRun ();_fa .Br =_aa .NewCT_TextLineBreak ();_b ._ed .EG_TextRun =append (_b ._ed .EG_TextRun ,_fa );};

// SetLevel sets the level of indentation of a paragraph.
func (_fg ParagraphProperties )SetLevel (idx int32 ){_fg ._dff .LvlAttr =_a .Int32 (idx )};

// SetSize sets the font size of the run text
func (_gg RunProperties )SetSize (sz _da .Distance ){_gg ._cge .SzAttr =_a .Int32 (int32 (sz /_da .HundredthPoint ));};

// X returns the inner wrapped XML type.
func (_gc LineProperties )X ()*_aa .CT_LineProperties {return _gc ._g };

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_aa .CT_TextParagraph )Paragraph {return Paragraph {x }};

// Properties returns the paragraph properties.
func (_db Paragraph )Properties ()ParagraphProperties {if _db ._ed .PPr ==nil {_db ._ed .PPr =_aa .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_db ._ed .PPr );};

// X returns the inner wrapped XML type.
func (_bfc Run )X ()*_aa .EG_TextRun {return _bfc ._ga };

// SetBold controls the bolding of a run.
func (_cga RunProperties )SetBold (b bool ){_cga ._cge .BAttr =_a .Bool (b )};

// SetJoin sets the line join style.
func (_f LineProperties )SetJoin (e LineJoin ){_f ._g .Round =nil ;_f ._g .Miter =nil ;_f ._g .Bevel =nil ;switch e {case LineJoinRound :_f ._g .Round =_aa .NewCT_LineJoinRound ();case LineJoinBevel :_f ._g .Bevel =_aa .NewCT_LineJoinBevel ();case LineJoinMiter :_f ._g .Miter =_aa .NewCT_LineJoinMiterProperties ();};};

// SetGeometry sets the shape type of the shape
func (_bg ShapeProperties )SetGeometry (g _aa .ST_ShapeType ){if _bg ._fc .PrstGeom ==nil {_bg ._fc .PrstGeom =_aa .NewCT_PresetGeometry2D ();};_bg ._fc .PrstGeom .PrstAttr =g ;};

// SetBulletChar sets the bullet character for the paragraph.
func (_ag ParagraphProperties )SetBulletChar (c string ){if c ==""{_ag ._dff .BuChar =nil ;}else {_ag ._dff .BuChar =_aa .NewCT_TextCharBullet ();_ag ._dff .BuChar .CharAttr =c ;};};

// SetSize sets the width and height of the shape.
func (_ff ShapeProperties )SetSize (w ,h _da .Distance ){_ff .SetWidth (w );_ff .SetHeight (h )};

// X returns the inner wrapped XML type.
func (_cbf ShapeProperties )X ()*_aa .CT_ShapeProperties {return _cbf ._fc };

// SetFlipVertical controls if the shape is flipped vertically.
func (_gb ShapeProperties )SetFlipVertical (b bool ){_gb .ensureXfrm ();if !b {_gb ._fc .Xfrm .FlipVAttr =nil ;}else {_gb ._fc .Xfrm .FlipVAttr =_a .Bool (true );};};

// SetSolidFill controls the text color of a run.
func (_cc RunProperties )SetSolidFill (c _c .Color ){_cc ._cge .NoFill =nil ;_cc ._cge .BlipFill =nil ;_cc ._cge .GradFill =nil ;_cc ._cge .GrpFill =nil ;_cc ._cge .PattFill =nil ;_cc ._cge .SolidFill =_aa .NewCT_SolidColorFillProperties ();_cc ._cge .SolidFill .SrgbClr =_aa .NewCT_SRgbColor ();_cc ._cge .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// X returns the inner wrapped XML type.
func (_cg Paragraph )X ()*_aa .CT_TextParagraph {return _cg ._ed };

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_e LineProperties )SetWidth (w _da .Distance ){_e ._g .WAttr =_a .Int32 (int32 (w /_da .EMU ))};func MakeShapeProperties (x *_aa .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};func (_ba ShapeProperties )SetSolidFill (c _c .Color ){_ba .clearFill ();_ba ._fc .SolidFill =_aa .NewCT_SolidColorFillProperties ();_ba ._fc .SolidFill .SrgbClr =_aa .NewCT_SRgbColor ();_ba ._fc .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};func (_aac LineProperties )SetSolidFill (c _c .Color ){_aac .clearFill ();_aac ._g .SolidFill =_aa .NewCT_SolidColorFillProperties ();_aac ._g .SolidFill .SrgbClr =_aa .NewCT_SRgbColor ();_aac ._g .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// Paragraph is a paragraph within a document.
type Paragraph struct{_ed *_aa .CT_TextParagraph };func (_ca ShapeProperties )LineProperties ()LineProperties {if _ca ._fc .Ln ==nil {_ca ._fc .Ln =_aa .NewCT_LineProperties ();};return LineProperties {_ca ._fc .Ln };};

// AddRun adds a new run to a paragraph.
func (_cb Paragraph )AddRun ()Run {_df :=MakeRun (_aa .NewEG_TextRun ());_cb ._ed .EG_TextRun =append (_cb ._ed .EG_TextRun ,_df .X ());return _df ;};

// SetHeight sets the height of the shape.
func (_daf ShapeProperties )SetHeight (h _da .Distance ){_daf .ensureXfrm ();if _daf ._fc .Xfrm .Ext ==nil {_daf ._fc .Xfrm .Ext =_aa .NewCT_PositiveSize2D ();};_daf ._fc .Xfrm .Ext .CyAttr =int64 (h /_da .EMU );};

// SetWidth sets the width of the shape.
func (_ebd ShapeProperties )SetWidth (w _da .Distance ){_ebd .ensureXfrm ();if _ebd ._fc .Xfrm .Ext ==nil {_ebd ._fc .Xfrm .Ext =_aa .NewCT_PositiveSize2D ();};_ebd ._fc .Xfrm .Ext .CxAttr =int64 (w /_da .EMU );};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_aa .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// RunProperties controls the run properties.
type RunProperties struct{_cge *_aa .CT_TextCharacterProperties ;};

// Run is a run within a paragraph.
type Run struct{_ga *_aa .EG_TextRun };

// GetPosition gets the position of the shape in EMU.
func (_cgg ShapeProperties )GetPosition ()(int64 ,int64 ){_cgg .ensureXfrm ();if _cgg ._fc .Xfrm .Off ==nil {_cgg ._fc .Xfrm .Off =_aa .NewCT_Point2D ();};return *_cgg ._fc .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_cgg ._fc .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// X returns the inner wrapped XML type.
func (_bf ParagraphProperties )X ()*_aa .CT_TextParagraphProperties {return _bf ._dff };

// Properties returns the run's properties.
func (_ge Run )Properties ()RunProperties {if _ge ._ga .R ==nil {_ge ._ga .R =_aa .NewCT_RegularTextRun ();};if _ge ._ga .R .RPr ==nil {_ge ._ga .R .RPr =_aa .NewCT_TextCharacterProperties ();};return RunProperties {_ge ._ga .R .RPr };};func (_gf LineProperties )clearFill (){_gf ._g .NoFill =nil ;_gf ._g .GradFill =nil ;_gf ._g .SolidFill =nil ;_gf ._g .PattFill =nil ;};

// SetText sets the run's text contents.
func (_dad Run )SetText (s string ){_dad ._ga .Br =nil ;_dad ._ga .Fld =nil ;if _dad ._ga .R ==nil {_dad ._ga .R =_aa .NewCT_RegularTextRun ();};_dad ._ga .R .T =s ;};