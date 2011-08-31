package wiki

import (
	"appengine"
	"appengine/urlfetch"

	"http"
	"regexp"
	"template"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}

	body := resp.Body
	var allData []byte
	allData = make([]byte, resp.ContentLength)
	body.Read(allData)

	firstLine := getFirstLine(string(allData))
	url := getUrl(string(allData))

	// TODO - check if firstLine is empty, and if so, log it
	if firstLine == "" {
		c.Errorf("Saw empty first line for data: %s", allData)
	}

	data := map[string] string {
		"firstLine" : firstLine,
		"url" : url,
	}


	err2 := mainPageTemplate.Execute(w, data)
	if err2 != nil {
		http.Error(w, err2.String(), http.StatusInternalServerError)
	}
}

func getFirstLine(stringToRegex string) string {
	getFirstParagraphRegex := regexp.MustCompile("<p><b>([^.]*).")
	firstLine := getFirstParagraphRegex.FindString(stringToRegex)

	return removeString(firstLine, "<([^<>]*)>")
}

func getUrl(stringToRegex string) string {
	getUrlRegex := regexp.MustCompile("Retrieved from \"<a href=\".*\">.*</a>\"")
	urlInfo := getUrlRegex.FindString(stringToRegex)

	findLinkText := regexp.MustCompile(">http://.*<")
	url := findLinkText.FindString(urlInfo)
	url = removeString(url, "<")
	url = removeString(url, ">")
	return url
}

func removeString(stringToRemove, strRegex string) string {
	actualRegex := regexp.MustCompile(strRegex)
	return actualRegex.ReplaceAllString(stringToRemove, "")
}

const testData = `
<meta http-equiv="Content-Style-Type" content="text/css" />
<meta name="generator" content="MediaWiki 1.17wmf1" />
<link rel="alternate" type="application/x-wiki" title="Edit this page" href="/w/index.php?title=Alton,_Staffordshire&amp;action=edit" />
<link rel="edit" title="Edit this page" href="/w/index.php?title=Alton,_Staffordshire&amp;action=edit" />
<link rel="apple-touch-icon" href="http://en.wikipedia.org/apple-touch-icon.png" />
<link rel="shortcut icon" href="/favicon.ico" />
<link rel="search" type="application/opensearchdescription+xml" href="/w/opensearch_desc.php" title="Wikipedia (en)" />
<link rel="EditURI" type="application/rsd+xml" href="http://en.wikipedia.org/w/api.php?action=rsd" />
<link rel="copyright" href="http://creativecommons.org/licenses/by-sa/3.0/" />
<link rel="alternate" type="application/atom+xml" title="Wikipedia Atom feed" href="/w/index.php?title=Special:RecentChanges&amp;feed=atom" />
<link rel="stylesheet" href="http://bits.wikimedia.org/en.wikipedia.org/load.php?debug=false&amp;lang=en&amp;modules=mediawiki%21legacy%21commonPrint%7Cmediawiki%21legacy%21shared%7Cskins%21vector&amp;only=styles&amp;skin=vector" type="text/css" media="all" />
<meta name="ResourceLoaderDynamicStyles" content="" /><link rel="stylesheet" href="http://bits.wikimedia.org/en.wikipedia.org/load.php?debug=false&amp;lang=en&amp;modules=site&amp;only=styles&amp;skin=vector" type="text/css" media="all" />
<style type="text/css" media="all">a.new,#quickbar a.new{color:#ba0000}

/* cache key: enwiki:resourceloader:filter:minify-css:5:f2a9127573a22335c2a9102b208c73e7 */</style>
<script src="http://bits.wikimedia.org/en.wikipedia.org/load.php?debug=false&amp;lang=en&amp;modules=startup&amp;only=scripts&amp;skin=vector" type="text/javascript"></script>
<script type="text/javascript">if ( window.mediaWiki ) {
	mediaWiki.config.set({"wgCanonicalNamespace": "", "wgCanonicalSpecialPageName": false, "wgNamespaceNumber": 0, "wgPageName": "Alton,_Staffordshire", "wgTitle": "Alton, Staffordshire", "wgAction": "view", "wgArticleId": 153690, "wgIsArticle": true, "wgUserName": null, "wgUserGroups": ["*"], "wgCurRevisionId": 425630865, "wgCategories": ["Articles with OS grid coordinates", "All articles with unsourced statements", "Articles with unsourced statements from January 2009", "Villages in Staffordshire", "Towns and villages of the Peak District", "Staffordshire Moorlands", "Staffordshire geography stubs"], "wgBreakFrames": false, "wgRestrictionEdit": [], "wgRestrictionMove": [], "wgSearchNamespaces": [0], "wgFlaggedRevsParams": {"tags": {"status": {"levels": 1, "quality": 2, "pristine": 3}}}, "wgStableRevisionId": null, "wgRevContents": {"error": "Unable to get content.", "waiting": "Waiting for content"}, "wgVectorEnabledModules": {"collapsiblenav": true, "collapsibletabs": true, "editwarning": true, "expandablesearch": false, "footercleanup": false, "sectioneditlinks": false, "simplesearch": true, "experiments": true}, "wgWikiEditorEnabledModules": {"toolbar": true, "dialogs": true, "templateEditor": false, "templates": false, "addMediaWizard": false, "preview": false, "previewDialog": false, "publish": false, "toc": false}, "wgTrackingToken": "e387a2df73041bffa8dd2fa3e39a2573", "wikilove-recipient": "", "wikilove-edittoken": "+\\", "wikilove-anon": 0, "mbEditToken": "+\\", "Geo": {"city": "", "country": ""}, "wgNoticeProject": "wikipedia"});
}
</script>

<!--[if lt IE 7]><style type="text/css">body{behavior:url("/w/skins-1.17/vector/csshover.min.htc")}</style><![endif]--></head>
<body class="mediawiki ltr ns-0 ns-subject page-Alton_Staffordshire skin-vector">
		<div id="mw-page-base" class="noprint"></div>
		<div id="mw-head-base" class="noprint"></div>
		<!-- content -->
		<div id="content">
			<a id="top"></a>
			<div id="mw-js-message" style="display:none;"></div>
						<!-- sitenotice -->
			<div id="siteNotice"><!-- centralNotice loads here --></div>
			<!-- /sitenotice -->
						<!-- firstHeading -->
			<h1 id="firstHeading" class="firstHeading">Alton, Staffordshire</h1>
			<!-- /firstHeading -->
			<!-- bodyContent -->
			<div id="bodyContent">
				<!-- tagline -->
				<div id="siteSub">From Wikipedia, the free encyclopedia</div>
				<!-- /tagline -->
				<!-- subtitle -->
				<div id="contentSub"></div>
				<!-- /subtitle -->
																<!-- jumpto -->
				<div id="jump-to-nav">
					Jump to: <a href="#mw-head">navigation</a>,
					<a href="#p-search">search</a>
				</div>
				<!-- /jumpto -->
								<!-- bodytext -->
				<div class="dablink">For other places with the same name, see <a href="/wiki/Alton_(disambiguation)" title="Alton (disambiguation)" class="mw-redirect">Alton (disambiguation)</a>.</div>
<p><span style="font-size: small;"><span id="coordinates"><a href="/wiki/Geographic_coordinate_system" title="Geographic coordinate system">Coordinates</a>: <span class="plainlinks nourlexpansion"><a href="http://toolserver.org/~geohack/geohack.php?pagename=Alton,_Staffordshire&amp;params=52.977_N_-1.890_E_region:GB_type:city" class="external text" rel="nofollow"><span class="geo-nondefault"><span class="geo-dms" title="Maps, aerial photos, and other data for this location"><span class="latitude">52°58′37″N</span> <span class="longitude">1°53′24″W</span></span></span><span class="geo-multi-punct">﻿ / ﻿</span><span class="geo-default"><span class="geo-dec" title="Maps, aerial photos, and other data for this location">52.977°N 1.890°W</span><span style="display:none">﻿ / <span class="geo">52.977; -1.890</span></span></span></a></span></span></span></p>
<table class="infobox geography vcard" style="width: 23em">
<tr class="mergedtoprow">
<td style="FONT-SIZE: 1.25em; width: 100%; white-space: nowrap; text-align: center" align="center" colspan="2"><span class="fn org"><b>Alton</b></span></td>
</tr>
<tr class="mergedtoprow">
<td style="text-align: center; padding: 0.7em 0.8em 0.7em 0.8em" align="center" colspan="2">
<center>
<div class="center">
<div style="width:240px; float: none; clear: both; margin-left: auto; margin-right: auto">
<div>
<div style="position: relative;"><a href="/wiki/File:Staffordshire_UK_location_map.svg" class="image" title="Alton is located in Staffordshire"><img alt="Alton is located in Staffordshire" src="http://upload.wikimedia.org/wikipedia/commons/thumb/0/09/Staffordshire_UK_location_map.svg/240px-Staffordshire_UK_location_map.svg.png" width="240" height="306" /></a>
<div style="position: absolute; z-index: 2; top: 32.1%; left: 55.5%; height: 0; width: 0; margin: 0; padding: 0;">
<div style="position: relative; text-align: center; left: -3px; top: -3px; width: 6px; font-size: 6px;"><img alt="" src="http://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Red_pog.svg/6px-Red_pog.svg.png" width="6" height="6" /></div>
<div style="font-size:&#160;%; line-height: 110%; z-index:90; position: relative; top: -1.5em; width: 6em; left: 0.5em; text-align: left;"><span style="padding: 1px;">Alton</span></div>
</div>
</div>
<div style="font-size: 90%; padding-top:3px;"></div>
</div>
</div>
</div>
<br />
<i><small><img alt="" src="http://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/Red_pog.svg/6px-Red_pog.svg.png" width="6" height="6" />&#160;Alton shown within <a href="/wiki/Staffordshire" title="Staffordshire">Staffordshire</a></small></i></center>
</td>
</tr>
<tr class="mergedtoprow">
<th><a href="/wiki/Ordnance_Survey_National_Grid" title="Ordnance Survey National Grid">OS&#160;grid&#160;reference</a></th>
<td><span class="plainlinks nourlexpansion"><span style="white-space: nowrap"><a href="http://toolserver.org/~rhaworth/os/coor_g.php?pagename=Alton,_Staffordshire&amp;params=SK073422_region%3AGB_scale%3A25000" class="external text" rel="nofollow">SK073422</a></span></span></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Districts_of_England" title="Districts of England">District</a></th>
<td><a href="/wiki/Staffordshire_Moorlands" title="Staffordshire Moorlands">Staffordshire Moorlands</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Metropolitan_and_non-metropolitan_counties_of_England" title="Metropolitan and non-metropolitan counties of England">Shire&#160;county</a></th>
<td><a href="/wiki/Staffordshire" title="Staffordshire">Staffordshire</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Regions_of_England" title="Regions of England">Region</a></th>
<td><a href="/wiki/West_Midlands_(region)" title="West Midlands (region)">West Midlands</a></td>
</tr>
<tr class="mergedrow adr">
<th><a href="/wiki/Countries_of_the_United_Kingdom" title="Countries of the United Kingdom">Country</a></th>
<td class="country-name"><a href="/wiki/England" title="England">England</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/List_of_sovereign_states" title="List of sovereign states">Sovereign&#160;state</a></th>
<td><a href="/wiki/United_Kingdom" title="United Kingdom">United Kingdom</a></td>
</tr>
<tr class="mergedtoprow">
<th><a href="/wiki/Post_town" title="Post town">Post town</a></th>
<td><span style="FONT-SIZE: 80%"><span style="text-transform:uppercase;"><a href="/wiki/Stoke-on-Trent" title="Stoke-on-Trent">Stoke-on-Trent</a></span></span></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Postcodes_in_the_United_Kingdom" title="Postcodes in the United Kingdom">Postcode&#160;district</a></th>
<td><span style="FONT-SIZE: 100%"><a href="/wiki/ST_postcode_area" title="ST postcode area">ST10</a></span></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Telephone_numbers_in_the_United_Kingdom" title="Telephone numbers in the United Kingdom">Dialling&#160;code</a></th>
<td><span style="FONT-SIZE: 100%">01538</span></td>
</tr>
<tr class="mergedtoprow">
<th><a href="/wiki/List_of_law_enforcement_agencies_in_the_United_Kingdom" title="List of law enforcement agencies in the United Kingdom">Police</a></th>
<td><a href="/wiki/Staffordshire_Police" title="Staffordshire Police">Staffordshire</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Fire_service_in_the_United_Kingdom" title="Fire service in the United Kingdom" class="mw-redirect">Fire</a></th>
<td><a href="/wiki/Staffordshire_Fire_and_Rescue_Service" title="Staffordshire Fire and Rescue Service">Staffordshire</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/Emergency_medical_services_in_the_United_Kingdom" title="Emergency medical services in the United Kingdom">Ambulance</a></th>
<td><a href="/wiki/West_Midlands_Ambulance_Service" title="West Midlands Ambulance Service">West Midlands</a></td>
</tr>
<tr class="mergedtoprow">
<th><a href="/wiki/Members_of_the_European_Parliament_for_the_United_Kingdom_2009%E2%80%932014" title="Members of the European Parliament for the United Kingdom 2009–2014">EU&#160;Parliament</a></th>
<td><a href="/wiki/West_Midlands_(European_Parliament_constituency)" title="West Midlands (European Parliament constituency)">West Midlands</a></td>
</tr>
<tr class="mergedrow">
<th><a href="/wiki/List_of_United_Kingdom_Parliament_constituencies" title="List of United Kingdom Parliament constituencies">UK&#160;Parliament</a></th>
<td><a href="/wiki/Staffordshire_Moorlands_(UK_Parliament_constituency)" title="Staffordshire Moorlands (UK Parliament constituency)">Staffordshire Moorlands</a></td>
</tr>
<tr class="mergedtoprow">
<td style="text-align: center" align="center" colspan="2"><small>List of places: <a href="/wiki/List_of_United_Kingdom_locations" title="List of United Kingdom locations">UK</a>&#160;•  <a href="/wiki/List_of_places_in_England" title="List of places in England">England</a>&#160;•  <a href="/wiki/List_of_places_in_Staffordshire" title="List of places in Staffordshire">Staffordshire</a></small></td>
</tr>
</table>
<p><b>Alton</b> is a <a href="/wiki/Village" title="Village">village</a> in the <a href="/wiki/County" title="County">county</a> of <a href="/wiki/Staffordshire" title="Staffordshire">Staffordshire</a>, <a href="/wiki/England" title="England">England</a>. It is noted for the theme park <a href="/wiki/Alton_Towers" title="Alton Towers">Alton Towers</a>, built around the site of Alton Mansion (also named Alton Towers), which was owned by the <a href="/wiki/Earl_of_Shrewsbury" title="Earl of Shrewsbury">Earls of Shrewsbury</a> and designed by <a href="/wiki/Augustus_Pugin" title="Augustus Pugin" class="mw-redirect">Augustus Pugin</a>.</p>
<div class="thumb tleft">
<div class="thumbinner" style="width:222px;"><a href="/wiki/File:Alton_Staffordshire_Railway_Station.jpg" class="image"><img alt="" src="http://upload.wikimedia.org/wikipedia/commons/thumb/0/0d/Alton_Staffordshire_Railway_Station.jpg/220px-Alton_Staffordshire_Railway_Station.jpg" width="220" height="165" class="thumbimage" /></a>
<div class="thumbcaption">
<div class="magnify"><a href="/wiki/File:Alton_Staffordshire_Railway_Station.jpg" class="internal" title="Enlarge"><img src="http://bits.wikimedia.org/skins-1.17/common/images/magnify-clip.png" width="15" height="11" alt="" /></a></div>
The disused Alton railway station.</div>
</div>
</div>
<p>The village is located on the eastern side of the <a href="/wiki/River_Churnet" title="River Churnet">Churnet</a> valley. It is mentioned in the <a href="/wiki/Domesday_Book" title="Domesday Book">Domesday Book</a>, and contains numerous buildings of architectural interest; the Round-House, <a href="/wiki/Alton_Castle" title="Alton Castle">Alton Castle</a> (now a <a href="/wiki/Catholic" title="Catholic">Catholic</a> youth <a href="/wiki/Retreat_(spiritual)" title="Retreat (spiritual)">retreat</a> centre), St Peter's Church, The Malt House, St John's Church and of course Alton Towers.</p>
<p>Alton was served by <a href="/wiki/Alton_railway_station,_Staffordshire" title="Alton railway station, Staffordshire" class="mw-redirect">Alton railway station</a> which was opened by the <a href="/wiki/North_Staffordshire_Railway" title="North Staffordshire Railway">North Staffordshire Railway</a> on July 13, 1849 and closed in the 1960s.</p>
<p>The <a href="/wiki/Chained_oak" title="Chained oak">chained oak</a> in Alton has been made famous by the ride Hex at <a href="/wiki/Alton_Towers" title="Alton Towers">Alton Towers</a> and the legend involving the <a href="/wiki/Earl_of_Shrewsbury" title="Earl of Shrewsbury">Earl of Shrewsbury</a>.</p>
<p>The village was home to seven public houses, including 'The Talbot', 'The Bulls Head', ' The Royal Oak', 'The Bridge House', 'The White Hart', 'The Blacksmiths Arms' and 'The Lord Shrewsbury' (formerly The Wild Duck, now wrongly named. Should have been the 'Earl of Shrewsbury'). The Talbot and The Lord Shrewsbury closed in 2008.</p>
<p>For those who believe in ghosts, Alton is also considered to be among the most haunted villages in Staffordshire. In particular, the ghost of a figure wearing a top hat and riding a horse has allegedly been sighted numerous times wandering through the fields around the village.<sup class="Template-Fact" style="white-space:nowrap;">[<i><a href="/wiki/Wikipedia:Citation_needed" title="Wikipedia:Citation needed"><span title="This claim needs references to reliable sources from January 2009">citation needed</span></a></i>]</sup> During the lifetime of <a href="/wiki/Pugin" title="Pugin" class="mw-redirect">Pugin</a> the village was known as Alverton.</p>
<h2><span class="editsection">[<a href="/w/index.php?title=Alton,_Staffordshire&amp;action=edit&amp;section=1" title="Edit section: External links">edit</a>]</span> <span class="mw-headline" id="External_links">External links</span></h2>
<table class="metadata mbox-small plainlinks" style="border:1px solid #aaa; background-color:#f9f9f9;">
<tr>
<td class="mbox-image"><img alt="" src="http://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Commons-logo.svg/30px-Commons-logo.svg.png" width="30" height="40" /></td>
<td class="mbox-text" style="">Wikimedia Commons has media related to: <i><b><a href="http://commons.wikimedia.org/wiki/Category:Alton,_Staffordshire" class="extiw" title="commons:Category:Alton, Staffordshire">Alton, Staffordshire</a></b></i></td>
</tr>
</table>
<ul>
<li><a href="http://www.crsbi.ac.uk/ed/st/alton/index.htm" class="external text" rel="nofollow">Information on Saint Peter's church</a></li>
<li><a href="http://www.wishful-thinking.org.uk/genuki/STS/Alton/StPeter.html" class="external text" rel="nofollow">Saint Peter's church</a></li>
<li><a href="http://www.francisfrith.com/search/england/staffordshire/alton/alton.htm" class="external text" rel="nofollow">Photographs of Alton</a></li>
<li><span class="citation web"><a href="http://www.alton-staffordshire.co.uk/default_frame.php?url=history.php" class="external text" rel="nofollow">"Alton History"</a>. Alton in Staffordshire<span class="printonly">. <a href="http://www.alton-staffordshire.co.uk/default_frame.php?url=history.php" class="external free" rel="nofollow">http://www.alton-staffordshire.co.uk/default_frame.php?url=history.php</a></span><span class="reference-accessdate">. Retrieved 2007-11-29</span>.</span><span class="Z3988" title="ctx_ver=Z39.88-2004&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook&amp;rft.genre=bookitem&amp;rft.btitle=Alton+History&amp;rft.atitle=&amp;rft.pub=Alton+in+Staffordshire&amp;rft_id=http%3A%2F%2Fwww.alton-staffordshire.co.uk%2Fdefault_frame.php%3Furl%3Dhistory.php&amp;rfr_id=info:sid/en.wikipedia.org:Alton,_Staffordshire"><span style="display: none;">&#160;</span></span></li>
</ul>
<p><br /></p>
<table class="metadata plainlinks stub" style="background: transparent;">
<tr>
<td><a href="/wiki/File:EnglandStaffordshire.svg" class="image"><img alt="Stub icon" src="http://upload.wikimedia.org/wikipedia/commons/thumb/1/1a/EnglandStaffordshire.svg/25px-EnglandStaffordshire.svg.png" width="25" height="31" /></a></td>
<td><i>This <a href="/wiki/Staffordshire" title="Staffordshire">Staffordshire</a> location article is a <a href="/wiki/Wikipedia:Stub" title="Wikipedia:Stub">stub</a>. You can help Wikipedia by <a href="http://en.wikipedia.org/w/index.php?title=Alton,_Staffordshire&amp;action=edit" class="external text" rel="nofollow">expanding it</a>.</i><span class="noprint plainlinks navbar" style="position:absolute; right:15px; font-size:smaller; display:none;"><span style="white-space:nowrap;word-spacing:-.12em;"><a href="/wiki/Template:Staffordshire-geo-stub" title="Template:Staffordshire-geo-stub"><span style="" title="View this template">v</span></a> <span style=""><b>·</b></span> <a href="/wiki/Template_talk:Staffordshire-geo-stub" title="Template talk:Staffordshire-geo-stub"><span style="" title="Discuss this template">d</span></a> <span style=""><b>·</b></span> <a href="http://en.wikipedia.org/w/index.php?title=Template:Staffordshire-geo-stub&amp;action=edit" class="external text" rel="nofollow"><span style="" title="Edit this template">e</span></a></span></span></td>
</tr>
</table>


<!-- 
NewPP limit report
Preprocessor node count: 2958/1000000
Post-expand include size: 32957/2048000 bytes
Template argument size: 7395/2048000 bytes
Expensive parser function count: 1/500
-->

<!-- Saved in parser cache with key enwiki:pcache:idhash:153690-0!*!0!!en!4 and timestamp 20110830101212 -->
<div class="printfooter">
Retrieved from "<a href="http://en.wikipedia.org/wiki/Alton,_Staffordshire">http://en.wikipedia.org/wiki/Alton,_Staffordshire</a>"</div>
				<!-- /bodytext -->
								<!-- catlinks -->
				<div id='catlinks' class='catlinks'><div id="mw-normal-catlinks"><a href="/wiki/Special:Categories" title="Special:Categories">Categories</a>: <span dir='ltr'><a href="/wiki/Category:Villages_in_Staffordshire" title="Category:Villages in Staffordshire">Villages in Staffordshire</a></span> | <span dir='ltr'><a href="/wiki/Category:Towns_and_villages_of_the_Peak_District" title="Category:Towns and villages of the Peak District">Towns and villages of the Peak District</a></span> | <span dir='ltr'><a href="/wiki/Category:Staffordshire_Moorlands" title="Category:Staffordshire Moorlands">Staffordshire Moorlands</a></span> | <span dir='ltr'><a href="/wiki/Category:Staffordshire_geography_stubs" title="Category:Staffordshire geography stubs">Staffordshire geography stubs</a></span></div><div id="mw-hidden-catlinks" class="mw-hidden-cats-hidden">Hidden categories: <span dir='ltr'><a href="/wiki/Category:Articles_with_OS_grid_coordinates" title="Category:Articles with OS grid coordinates">Articles with OS grid coordinates</a></span> | <span dir='ltr'><a href="/wiki/Category:All_articles_with_unsourced_statements" title="Category:All articles with unsourced statements">All articles with unsourced statements</a></span> | <span dir='ltr'><a href="/wiki/Category:Articles_with_unsourced_statements_from_January_2009" title="Category:Articles with unsourced statements from January 2009">Articles with unsourced statements from January 2009</a></span></div></div>	<!-- /catlinks -->
												<div class="visualClear"></div>
			</div>
			<!-- /bodyContent -->
		</div>
		<!-- /content -->
		<!-- header -->
		<div id="mw-head" class="noprint">
			
<!-- 0 -->
<div id="p-personal" class="">
	<h5>Personal tools</h5>
	<ul>
					<li  id="pt-login"><a href="/w/index.php?title=Special:UserLogin&amp;returnto=Alton,_Staffordshire" title="You are encouraged to log in; however, it is not mandatory. [o]" accesskey="o">Log in / create account</a></li>
			</ul>
</div>

<!-- /0 -->
			<div id="left-navigation">
				
<!-- 0 -->
<div id="p-namespaces" class="vectorTabs">
	<h5>Namespaces</h5>
	<ul>
					<li  id="ca-nstab-main" class="selected"><span><a href="/wiki/Alton,_Staffordshire"  title="View the content page [c]" accesskey="c">Article</a></span></li>
					<li  id="ca-talk"><span><a href="/wiki/Talk:Alton,_Staffordshire"  title="Discussion about the content page [t]" accesskey="t">Discussion</a></span></li>
			</ul>
</div>

<!-- /0 -->

<!-- 1 -->
<div id="p-variants" class="vectorMenu emptyPortlet">
		<h5><span>Variants</span><a href="#"></a></h5>
	<div class="menu">
		<ul>
					</ul>
	</div>
</div>

<!-- /1 -->
			</div>
			<div id="right-navigation">
				
<!-- 0 -->
<div id="p-views" class="vectorTabs">
	<h5>Views</h5>
	<ul>
					<li id="ca-view" class="selected"><span><a href="/wiki/Alton,_Staffordshire" >Read</a></span></li>
					<li id="ca-edit"><span><a href="/w/index.php?title=Alton,_Staffordshire&amp;action=edit"  title="You can edit this page. &#10;Please use the preview button before saving. [e]" accesskey="e">Edit</a></span></li>
					<li id="ca-history" class="collapsible "><span><a href="/w/index.php?title=Alton,_Staffordshire&amp;action=history"  title="Past versions of this page [h]" accesskey="h">View history</a></span></li>
			</ul>
</div>

<!-- /0 -->

<!-- 1 -->
<div id="p-cactions" class="vectorMenu emptyPortlet">
	<h5><span>Actions</span><a href="#"></a></h5>
	<div class="menu">
		<ul>
					</ul>
	</div>
</div>

<!-- /1 -->

<!-- 2 -->
<div id="p-search">
	<h5><label for="searchInput">Search</label></h5>
	<form action="/w/index.php" id="searchform">
		<input type='hidden' name="title" value="Special:Search"/>
				<div id="simpleSearch">
						<input id="searchInput" name="search" type="text"  title="Search Wikipedia [f]" accesskey="f"  value="" />
						<button id="searchButton" type='submit' name='button'  title="Search Wikipedia for this text"><img src="http://bits.wikimedia.org/skins-1.17/vector/images/search-ltr.png?301-3" alt="Search" /></button>
					</div>
			</form>
</div>

<!-- /2 -->
			</div>
		</div>
		<!-- /header -->
		<!-- panel -->
			<div id="mw-panel" class="noprint">
				<!-- logo -->
					<div id="p-logo"><a style="background-image: url(http://upload.wikimedia.org/wikipedia/en/b/bc/Wiki.png);" href="/wiki/Main_Page"  title="Visit the main page"></a></div>
				<!-- /logo -->
				
<!-- navigation -->
<div class="portal" id='p-navigation'>
	<h5>Navigation</h5>
	<div class="body">
				<ul>
					<li id="n-mainpage-description"><a href="/wiki/Main_Page" title="Visit the main page [z]" accesskey="z">Main page</a></li>
					<li id="n-contents"><a href="/wiki/Portal:Contents" title="Guides to browsing Wikipedia">Contents</a></li>
					<li id="n-featuredcontent"><a href="/wiki/Portal:Featured_content" title="Featured content – the best of Wikipedia">Featured content</a></li>
					<li id="n-currentevents"><a href="/wiki/Portal:Current_events" title="Find background information on current events">Current events</a></li>
					<li id="n-randompage"><a href="/wiki/Special:Random" title="Load a random article [x]" accesskey="x">Random article</a></li>
					<li id="n-sitesupport"><a href="http://wikimediafoundation.org/wiki/Special:Landingcheck?landing_page=WMFJA085&amp;language=en&amp;utm_source=donate&amp;utm_medium=sidebar&amp;utm_campaign=20101204SB002" title="Support us">Donate to Wikipedia</a></li>
				</ul>
			</div>
</div>

<!-- /navigation -->

<!-- SEARCH -->

<!-- /SEARCH -->

<!-- interaction -->
<div class="portal" id='p-interaction'>
	<h5>Interaction</h5>
	<div class="body">
				<ul>
					<li id="n-help"><a href="/wiki/Help:Contents" title="Guidance on how to use and edit Wikipedia">Help</a></li>
					<li id="n-aboutsite"><a href="/wiki/Wikipedia:About" title="Find out about Wikipedia">About Wikipedia</a></li>
					<li id="n-portal"><a href="/wiki/Wikipedia:Community_portal" title="About the project, what you can do, where to find things">Community portal</a></li>
					<li id="n-recentchanges"><a href="/wiki/Special:RecentChanges" title="The list of recent changes in the wiki [r]" accesskey="r">Recent changes</a></li>
					<li id="n-contact"><a href="/wiki/Wikipedia:Contact_us" title="How to contact Wikipedia">Contact Wikipedia</a></li>
				</ul>
			</div>
</div>

<!-- /interaction -->

<!-- TOOLBOX -->
<div class="portal" id="p-tb">
	<h5>Toolbox</h5>
	<div class="body">
		<ul>
					<li id="t-whatlinkshere"><a href="/wiki/Special:WhatLinksHere/Alton,_Staffordshire" title="List of all English Wikipedia pages containing links to this page [j]" accesskey="j">What links here</a></li>
						<li id="t-recentchangeslinked"><a href="/wiki/Special:RecentChangesLinked/Alton,_Staffordshire" title="Recent changes in pages linked from this page [k]" accesskey="k">Related changes</a></li>
																					<li id="t-upload"><a href="/wiki/Wikipedia:Upload" title="Upload files [u]" accesskey="u">Upload file</a></li>
											<li id="t-specialpages"><a href="/wiki/Special:SpecialPages" title="List of all special pages [q]" accesskey="q">Special pages</a></li>
											<li id="t-permalink"><a href="/w/index.php?title=Alton,_Staffordshire&amp;oldid=425630865" title="Permanent link to this revision of the page">Permanent link</a></li>
				<li id="t-cite"><a href="/w/index.php?title=Special:Cite&amp;page=Alton,_Staffordshire&amp;id=425630865" title="Information on how to cite this page">Cite this page</a></li>		</ul>
	</div>
</div>

<!-- /TOOLBOX -->

<!-- coll-print_export -->
<div class="portal" id='p-coll-print_export'>
	<h5>Print/export</h5>
	<div class="body">
				<ul id="collectionPortletList"><li id="coll-create_a_book"><a href="/w/index.php?title=Special:Book&amp;bookcmd=book_creator&amp;referer=Alton%2C+Staffordshire" title="Create a book or page collection" rel="nofollow">Create a book</a></li><li id="coll-download-as-rl"><a href="/w/index.php?title=Special:Book&amp;bookcmd=render_article&amp;arttitle=Alton%2C+Staffordshire&amp;oldid=425630865&amp;writer=rl" title="Download a PDF version of this wiki page" rel="nofollow">Download as PDF</a></li><li id="t-print"><a href="/w/index.php?title=Alton,_Staffordshire&amp;printable=yes" title="Printable version of this page [p]" accesskey="p">Printable version</a></li></ul>			</div>
</div>

<!-- /coll-print_export -->

<!-- LANGUAGES -->
<div class="portal" id="p-lang">
	<h5>Languages</h5>
	<div class="body">
		<ul>
					<li class="interwiki-nl"><a href="http://nl.wikipedia.org/wiki/Alton_(Staffordshire)" title="Alton (Staffordshire)">Nederlands</a></li>
					<li class="interwiki-pl"><a href="http://pl.wikipedia.org/wiki/Alton_(Staffordshire)" title="Alton (Staffordshire)">Polski</a></li>
				</ul>
	</div>
</div>

<!-- /LANGUAGES -->
			</div>
		<!-- /panel -->
		<!-- footer -->
		<div id="footer">
											<ul id="footer-info">
																	<li id="footer-info-lastmod"> This page was last modified on 24 April 2011 at 08:33.<br /></li>
																					<li id="footer-info-copyright">Text is available under the <a rel="license" href="http://en.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License">Creative Commons Attribution-ShareAlike License</a><a rel="license" href="http://creativecommons.org/licenses/by-sa/3.0/" style="display:none;"></a>;
additional terms may apply.
See <a href="http://wikimediafoundation.org/wiki/Terms_of_use">Terms of use</a> for details.<br/>
Wikipedia&reg; is a registered trademark of the <a href="http://www.wikimediafoundation.org/">Wikimedia Foundation, Inc.</a>, a non-profit organization.<br /></li><li class="noprint"><a class='internal' href="http://en.wikipedia.org/wiki/Wikipedia:Contact_us">Contact us</a></li>
															</ul>
															<ul id="footer-places">
																	<li id="footer-places-privacy"><a href="http://wikimediafoundation.org/wiki/Privacy_policy" title="wikimedia:Privacy policy">Privacy policy</a></li>
																					<li id="footer-places-about"><a href="/wiki/Wikipedia:About" title="Wikipedia:About">About Wikipedia</a></li>
																					<li id="footer-places-disclaimer"><a href="/wiki/Wikipedia:General_disclaimer" title="Wikipedia:General disclaimer">Disclaimers</a></li>
																					<li id="footer-places-mobileview"><a href='/w/index.php?title=Alton,_Staffordshire&amp;useformat=mobile'>Mobile view</a></li>
															</ul>
											<ul id="footer-icons" class="noprint">
					<li id="footer-copyrightico">
						<a href="http://wikimediafoundation.org/"><img src="http://bits.wikimedia.org/images/wikimedia-button.png" width="88" height="31" alt="Wikimedia Foundation"/></a>
					</li>
					<li id="footer-poweredbyico">
						<a href="http://www.mediawiki.org/"><img src="http://bits.wikimedia.org/skins-1.17/common/images/poweredby_mediawiki_88x31.png" alt="Powered by MediaWiki" width="88" height="31" /></a>
					</li>
				</ul>
						<div style="clear:both"></div>
		</div>
		<!-- /footer -->
		<script type="text/javascript">if ( window.mediaWiki ) {
	mediaWiki.loader.load(["mediawiki.legacy.wikibits", "mediawiki.util", "mediawiki.legacy.ajax", "mediawiki.legacy.mwsuggest", "ext.vector.collapsibleNav", "ext.vector.collapsibleTabs", "ext.vector.editWarning", "ext.vector.simpleSearch", "ext.UserBuckets", "ext.articleFeedback.startup"]);
	mediaWiki.loader.go();
}
</script>

<script src="/w/index.php?title=Special:BannerController&amp;cache=/cn.js&amp;301-3" type="text/javascript"></script>
<script src="http://bits.wikimedia.org/en.wikipedia.org/load.php?debug=false&amp;lang=en&amp;modules=site&amp;only=scripts&amp;skin=vector" type="text/javascript"></script>
<script type="text/javascript">if ( window.mediaWiki ) {
	mediaWiki.user.options.set({"ccmeonemails":0,"cols":80,"contextchars":50,"contextlines":5,"date":"default","diffonly":0,"disablemail":0,"disablesuggest":0,"editfont":"default","editondblclick":0,"editsection":1,"editsectiononrightclick":0,"enotifminoredits":0,"enotifrevealaddr":0,"enotifusertalkpages":1,"enotifwatchlistpages":0,"extendwatchlist":0,"externaldiff":0,"externaleditor":0,"fancysig":0,"forceeditsummary":0,"gender":"unknown","hideminor":0,"hidepatrolled":0,"highlightbroken":1,"imagesize":2,"justify":0,"math":1,"minordefault":0,"newpageshidepatrolled":0,"nocache":0,"noconvertlink":0,"norollbackdiff":0,"numberheadings":0,"previewonfirst":0,"previewontop":1,"quickbar":1,"rcdays":7,"rclimit":50,"rememberpassword":0,"rows":25,"searchlimit":20,"showhiddencats":false,"showjumplinks":1,"shownumberswatching":1,"showtoc":1,"showtoolbar":1,"skin":"vector","stubthreshold":0,"thumbsize":4,"underline":2,"uselivepreview":0,"usenewrc":0,"watchcreations":1,"watchdefault":0,"watchdeletion":0,
	"watchlistdays":"3","watchlisthideanons":0,"watchlisthidebots":0,"watchlisthideliu":0,"watchlisthideminor":0,"watchlisthideown":0,"watchlisthidepatrolled":0,"watchmoves":0,"wllimit":250,"flaggedrevssimpleui":1,"flaggedrevsstable":false,"flaggedrevseditdiffs":true,"flaggedrevsviewdiffs":false,"vector-simplesearch":1,"useeditwarning":1,"vector-collapsiblenav":1,"usebetatoolbar":1,"usebetatoolbar-cgd":1,"wikilove-enabled":1,"variant":"en","language":"en","searchNs0":true,"searchNs1":false,"searchNs2":false,"searchNs3":false,"searchNs4":false,"searchNs5":false,"searchNs6":false,"searchNs7":false,"searchNs8":false,"searchNs9":false,"searchNs10":false,"searchNs11":false,"searchNs12":false,"searchNs13":false,"searchNs14":false,"searchNs15":false,"searchNs100":false,"searchNs101":false,"searchNs108":false,"searchNs109":false});;mediaWiki.loader.state({"user.options":"ready"});
	
	/* cache key: enwiki:resourceloader:filter:minify-js:5:c183491fdc987ec95b8873a74ef2bb96 */
}
</script><script type="text/javascript" src="http://geoiplookup.wikimedia.org/"></script>		<!-- fixalpha -->
		<script type="text/javascript"> if ( window.isMSIE55 ) fixalpha(); </script>
		<!-- /fixalpha -->
		<!-- Served by srv270 in 0.920 secs. -->			</body>
</html>
`

var mainPageTemplate = template.MustParse(mainPage, nil)

const mainPage = `
<html>
  <head>
    <title>Wiki Made Easy</title>
    <link rel="stylesheet" href="/static/css/bootstrap-1.1.0.min.css">
    <script src="static/js/jquery.js" type="text/javascript"></script>
    <script src="static/js/custom.js" type="text/javascript"></script>
  </head>
  <body>
    <div class="container">
      <br /><br /><br /><br /><br /><br /><br />
      <h1>Did you know...</h1>
      <h1>{firstLine|html}</h1>
      <h1><a href="{url|html}">Learn more about this!</a></h1>
      <h1><a href="/">Learn something else!</a></h1>
      <br /><br /><br /><br /><br /><br /><br />
      <div>
        <a href="http://code.google.com/appengine/">
          <img src="/static/img/appengine-silver-120x30.gif" alt="Powered by Google App Engine" />
        </a>
        <a href="http://golang.org">
          <img src="/static/img/Golang.png" alt="Powered by Go" />
        </a>
      </div>
    </div>
  </body>
</html>
`

