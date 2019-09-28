// Code generated by qtc from "docs.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type DocsPage struct {
	BasePage
}

func (p *DocsPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(` - Documentation
`)
}

func (p *DocsPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *DocsPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *DocsPage) StreamPageMeta(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.BasePage.StreamPageMeta(qw422016)
	qw422016.N().S(`
`)
	streamsidebarStyle(qw422016)
	qw422016.N().S(`
<style>
.title {
    padding-top: 2rem;
}
</style>
`)
}

func (p *DocsPage) WritePageMeta(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageMeta(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *DocsPage) PageMeta() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageMeta(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamipsum(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec placerat turpis eu odio finibus, eget consequat tellus cursus. Etiam in condimentum justo. Ut condimentum nulla eget lobortis porta. Etiam erat tellus, placerat vel turpis non, pretium sagittis magna. Cras dignissim felis nec diam malesuada molestie. Vestibulum convallis elit at arcu dictum, nec venenatis ex commodo. Nulla nec hendrerit orci, ut ullamcorper risus. Morbi venenatis augue eu lacus pretium cursus. Etiam nec enim commodo, bibendum felis quis, sagittis ex.</p>
<p>Aenean semper risus ipsum, sit amet laoreet lacus blandit nec. Vivamus gravida, dolor non faucibus tempor, ex ipsum tristique dui, vel convallis massa dui cursus eros. Cras porttitor magna vestibulum, dignissim sapien ac, porta felis. In fringilla viverra ex, nec scelerisque sapien hendrerit sed. Mauris placerat eu nisi quis accumsan. Nam finibus quis augue in mollis. Mauris laoreet quam quam, non efficitur mauris venenatis at. Sed venenatis, sapien in convallis hendrerit, nibh tellus posuere leo, non cursus eros libero et libero. Mauris rutrum, lectus eu ornare porta, dolor velit accumsan tortor, ut ornare nunc lectus faucibus felis. Mauris feugiat dui et vehicula eleifend. Maecenas sit amet erat et magna iaculis finibus auctor eu sem. Aliquam id arcu sed augue imperdiet lobortis quis id risus.</p>
<p>Mauris et turpis orci. Sed faucibus leo sem, pharetra tincidunt nunc accumsan eget. Cras condimentum vitae enim ut auctor. Quisque eget felis mattis, fringilla orci luctus, malesuada erat. Etiam scelerisque nunc id nibh porta aliquam. Interdum et malesuada fames ac ante ipsum primis in faucibus. Mauris vel urna elementum, feugiat libero nec, malesuada neque. Nunc ut nisi ut magna bibendum accumsan. Aenean efficitur massa in nunc tempus rutrum. Vivamus scelerisque a odio eu mollis. Aenean in lectus in mauris mattis elementum eget semper justo. Nullam sagittis vitae urna ut tincidunt. Nullam aliquet mi nec mauris congue bibendum non eget augue. Pellentesque leo ante, elementum aliquet commodo sed, congue quis massa. Quisque dapibus est metus, vel tempor velit vulputate quis.</p>
<p>Sed quis velit vestibulum, aliquet elit vitae, gravida lacus. Etiam dapibus semper felis, id euismod ipsum tristique sed. Nunc sem sapien, placerat ut mi sit amet, tincidunt bibendum ligula. In tempus risus at nunc luctus imperdiet. Etiam volutpat, mi vitae mollis bibendum, mauris mauris ullamcorper orci, et scelerisque tellus felis quis purus. Donec purus lorem, suscipit pharetra sapien non, congue faucibus nulla. Phasellus rhoncus ex quis elementum cursus. Nullam dictum ipsum et neque pharetra pellentesque. Donec felis tortor, rutrum ac dui eget, ullamcorper tempor nibh.</p>
<p>Cras vel sodales purus, sit amet ultricies nibh. Cras vitae odio sollicitudin tortor dictum malesuada et at justo. Pellentesque rhoncus ante non felis auctor, non porta mauris pharetra. Donec iaculis mauris eu ornare condimentum. Nulla tempus consequat lectus, in maximus eros viverra a. Sed fringilla felis quis erat elementum accumsan. Aenean ullamcorper enim vel diam tempor bibendum. Donec eu risus non quam tincidunt congue ac sit amet leo. Nulla sed purus commodo, fermentum elit eu, cursus risus. Nulla gravida mi ut mi consequat, a tincidunt metus dignissim. Sed sed mollis lectus, id placerat nulla. Duis tincidunt ligula et odio tincidunt laoreet. Nam aliquet odio ac tellus sollicitudin, eget commodo purus pulvinar.</p>
`)
}

func writeipsum(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamipsum(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func ipsum() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writeipsum(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *DocsPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<div class="columns is-fullheight is-clipped">
    <div class="is-sidebar-menu" id="sidebar">
        <aside class="menu">
            <p class="menu-label">
                General
            </p>
            <ul class="menu-list">
                <li><a href="#commands">Commands</a></li>
            </ul>
            <p class="menu-label">
                Custom commands
            </p>
            <ul class="menu-list">
                <li><a href="#triggers">Triggers</a></li>
            </ul>
        </aside>
    </div>

    <div class="column is-main-content content" id="main">
        <h1>General</h1>

        <section id="commands" class="page">
            <h2 class="title">Commands</h2>

            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
        </section>

        <h1>Custom commands</h1>

        <section id="triggers" class="page">
            <h2 class="title">Triggers</h2>

            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
            `)
	streamipsum(qw422016)
	qw422016.N().S(`
        </section>
    </div>
</div>
`)
}

func (p *DocsPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *DocsPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *DocsPage) StreamPageScripts(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.BasePage.StreamPageScripts(qw422016)
	qw422016.N().S(`
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.slim.min.js" integrity="sha256-pasqAKBDmFT4eHoN2ndd6lN370kFiGUFyTiUHWhU7k8=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/gumshoe/5.1.1/gumshoe.polyfills.min.js" integrity="sha256-nxQlu6aVkhkHICQuYc1Ga85x2MX+FpvJYV1F+O8Awy8=" crossorigin="anonymous"></script>

<script>
$(function() {
    document.addEventListener("gumshoeActivate", function(event) {
        $(event.detail.link).addClass("is-active");
    }, false);
    document.addEventListener("gumshoeDeactivate", function(event) {
        $(event.detail.link).removeClass("is-active");
    }, false);

    var header = document.querySelector("#header");
    spy = new Gumshoe("#sidebar a", {
        navClass: "is-active",
        contentClass: "is-active",
        offset: function() {
            return header.getBoundingClientRect().height;
        },
        events: true
    });

    spy.setup();
    spy.detect();

    $("#main").scroll(function() {
        spy.detect();
    });
});
</script>
`)
}

func (p *DocsPage) WritePageScripts(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageScripts(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *DocsPage) PageScripts() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageScripts(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func streamdocLevelTag(qw422016 *qt422016.Writer, level string) {
	qw422016.N().S(`
`)
	switch level {
	case "everyone":
		qw422016.N().S(`
<span class="tag is-dark">Everyone</span>
`)
	case "subs":
		qw422016.N().S(`
<span class="tag is-success">Subs</span>
`)
	case "mods":
		qw422016.N().S(`
<span class="tag is-warning">Mods</span>
`)
	case "broadcaster":
		qw422016.N().S(`
<span class="tag is-danger">Broadcaster</span>
`)
	case "admin":
		qw422016.N().S(`
<span class="tag is-info">Admins</span>
`)
	default:
		qw422016.N().S(`
`)
	}
	qw422016.N().S(`
`)
}

func writedocLevelTag(qq422016 qtio422016.Writer, level string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamdocLevelTag(qw422016, level)
	qt422016.ReleaseWriter(qw422016)
}

func docLevelTag(level string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writedocLevelTag(qb422016, level)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
