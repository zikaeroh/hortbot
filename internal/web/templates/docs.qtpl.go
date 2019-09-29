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
h2.title {
    padding-top: 1rem;
}

dt {
    padding-top: 1rem;
}

dd {
    padding-top: 0.5rem;
    padding-bottom: 0.5rem;
}

.tag {
    margin-left: 1rem;
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
                <li><a href="#repeats">Repeats</a></li>
                <li><a href="#schedule">Schedule</a></li>
                <li><a href="#autoreplies">Autoreplies</a></li>
            </ul>

            <p class="menu-label">
                Fun
            </p>
            <ul class="menu-list">
                <li><a href="#quotes">Quotes</a></li>
            </ul>
        </aside>
    </div>

                    
    <div class="column is-main-content content" id="main">
        <h1 class="title">General</h1>

        <section id="commands" class="page">
            <h2 class="title">Commands</h2>

            <dl>
                `)
	streamcommand(qw422016, "!join", `Tells the bot to join your channel. Must be executed in the bot's channel.`, "everyone")
	qw422016.N().S(`
                `)
	streamcommand(qw422016, "!part", `Tells the bot to leave your channel.`, "everyone")
	qw422016.N().S(`
            </dl>
        </section>

        <h1 class="title">Custom commands</h1>

        <section id="triggers" class="page">
            <h2 class="title">Triggers</h2>

            <dl>
                `)
	streamcommand(qw422016,
		"!commands",
		`Links to the list of commands for the channel.`,
		"subs",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command add <name> <text>",
		`Creates a command <code>!name</code> that responds with the specified text.`,
		"mods",
		`By default, commands are available to subs. Using <code>adda</code> or <code>addm</code> instead of <code>add</code> will pre-restrict the command to all users or moderators, respectively.`,
		`Example: <code>!command add pan FOUND THE (_PARAMETER_CAPS_), HAVE YE?</code> &mdash; Adds a command called "pan".`,
		`Example: <code>!command adda useful Here's some useful info: example.org</code> &mdash; Adds a command available to all users immediately.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command delete <name>",
		`Deletes a command.`,
		"mods",
		`Example: <code>!command delete pan</code> &mdash; Deletes the command called "pan".`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command restrict <name> all|subs|mods|owner",
		`Restricts a command to a specific group.`,
		"mods",
		`Example: <code>!command restrict pan mods</code> &mdash; Restricts "pan" to moderators and above.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command editor <name>",
		`Gets the last editor of a command.`,
		"mods",
		`Example: <code>!command editor pan</code> &mdash; Gets the last editor of the "pan" command.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command count <name>",
		`Gets the number of times a command has been run.`,
		"mods",
		`Example: <code>!command count pan</code> &mdash; Gets the number of times the "pan" command have been used.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command rename <old> <new>",
		`Renames a command.`,
		"mods",
		`Example: <code>!command rename pan oldpan</code> &mdash; Renames the command "pan" to "oldpan".`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!command rename <name>",
		`Gets the response for a command.`,
		"mods",
		`Example: <code>!command get pan</code> &mdash; Gets the response for the "pan" command.`,
	)
	qw422016.N().S(`
            </dl>
        </section>

        <section id="repeats" class="page">
            <h2 class="title">Repeats</h2>
            
            <p>
                The repeat command sets up a command repetition. When enabled,
                the bot will repeat every X seconds so long as Y messages have
                passed.
            </p>

            <dl>
                `)
	streamcommand(qw422016,
		"!repeat add <name> <delay in seconds> [message difference]",
		`Sets a command to repeat, and enables it.`,
		"mods",
		`Example: <code>!repeat add discord 300 10</code> &mdash; Sets the command "discord" to repeat every 300 seconds if at least 10 messages have passed.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!repeat delete <name>",
		`Deletes a command's repeat info.`,
		"mods",
		`Example: <code>!repeat delete discord</code> &mdash; Stops repeating the "discord" command and deletes its repeat info.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!repeat on|off <name>",
		`Enables or disables a command's repetition.`,
		"mods",
		`Example: <code>!repeat on discord</code> &mdash; Enables repetition of the "discord" command.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!repeat list",
		`Lists command repetition info.`,
		"mods",
	)
	qw422016.N().S(`
            </dl>
        </section>

        <section id="schedule" class="page">
            <h2 class="title">Schedule</h2>

            <p>
                The schedule command sets up a command repetition via a <a href="https://en.wikipedia.org/wiki/Cron" target="_blank">cron expression</a>.
                Like repeated commands, a message difference can be specified.
            </p>

            <dl>
                `)
	streamcommand(qw422016,
		"!schedule add <name> <pattern> [message difference]",
		`Schedules a command, and enables it.`,
		"mods",
		`Example: <code>!schedule add discord *_5_*_*_*</code> &mdash; Schedules the command "discord" to at 5AM every day.`,
		`Example: <code>!schedule add discord hourly 10</code> &mdash; Schedules the command "discord" to run hourly if at least 10 messages have passed.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!schedule delete <name>",
		`Deletes a command's schedule.`,
		"mods",
		`Example: <code>!schedule delete discord</code> &mdash; Unschedules the "discord" command and deletes its schedule.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!schedule on|off <name>",
		`Enables or disables a command's repetition.`,
		"mods",
		`Example: <code>!schedule on discord</code> &mdash; Enables the schedule of the "discord" command.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!schedule list",
		`Lists command schedules.`,
		"mods",
	)
	qw422016.N().S(`
            </dl>
        </section>

        <section id="autoreplies" class="page">
            <h2 class="title">Autoreplies</h2>

            <p>
                Autoreplies are like custom commands, but are run when a message matches a pattern.
            </p>

            <dl>
                `)
	streamcommand(qw422016,
		"!autoreply add <pattern> <response>",
		`Adds an autoreply which will respond with the provided response when a message matches the pattern.`,
		"mods",
		`In the pattern, spaces should be replaced with underscores.`,
		`Example: <code>!autoreply add *what*game* This is (_GAME_).</code> &mdash; Adds an autoreply that will reply with the current game if a message matches the pattern "*what*game".`,
		`Example: <code>!autoreply add REGEX:^too_many_[^_]+$ TOO MANY COOKS (_REGULARS_ONLY_)</code> &mdash; Adds an autoreply which uses a raw regex pattern.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!autoreply delete <num>",
		`Removes an autoreply.`,
		"mods",
		`Note that deleting an autoreply that isn't the last does not shift the numbers down. Use <code>!autoreply compact</code> to do this.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!autoreply editresponse <num> <response>",
		`Edits an autoreply's response.`,
		"mods",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!autoreply editpattern <num> <pattern>",
		`Edits an autoreply's pattern.`,
		"mods",
		`In the pattern, spaces should be replaced with underscores.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!autoreply compact <num>",
		`Compacts autoreplies "num" and higher. This is useful after removing an autoreply in the middle of the list.`,
		"mods",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!autoreply list",
		`Links to the list of autoreplies for the channel.`,
		"mods",
	)
	qw422016.N().S(`
            </dl>
        </section>

        <h1 class="title">Fun</h1>

        <section id="quotes" class="page">
            <h2 class="title">Quotes</h2>

            <dl>
                `)
	streamcommand(qw422016, "!quote", `Gets a random quote.`, "subs")
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote add <quote>",
		`Adds a quote.`,
		"mods",
		`Example: <code>!quote add "This is a quote!"</code> &mdash; Adds a the quote "This is a quote!".`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote delete <num>",
		`Removes a quote.`,
		"mods",
		`Note that deleting a quote that isn't the last does not shift the numbers down. Use <code>!quote compact</code> to do this.`,
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote get <num>",
		`Gets a quote.`,
		"subs",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote random",
		`Gets a random quote.`,
		"subs",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote getindex <exact quote>",
		`Returns the number of the exact quote specified.`,
		"subs",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote search <phrase>",
		`Searches all quotes for a phrase.`,
		"mods",
	)
	qw422016.N().S(`
                `)
	streamcommand(qw422016,
		"!quote compact <num>",
		`Compacts quotes "num" and higher. This is useful after removing a quote in the middle of the list.`,
		"mods",
	)
	qw422016.N().S(`
            </dl>
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
        // offset: function() {
        //     return header.getBoundingClientRect().height;
        // },
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

func streamcommand(qw422016 *qt422016.Writer, name, desc, level string, extra ...string) {
	qw422016.N().S(`
<dt><code>`)
	qw422016.E().S(name)
	qw422016.N().S(`</code></dt>
<dd>
    <p>
        `)
	qw422016.N().S(desc)
	qw422016.N().S(`
        `)
	streamdocLevelTag(qw422016, level)
	qw422016.N().S(`
    </p>
    `)
	for _, ex := range extra {
		qw422016.N().S(`
    <p>
        `)
		qw422016.N().S(ex)
		qw422016.N().S(`
    </p>
    `)
	}
	qw422016.N().S(`
</dd>
`)
}

func writecommand(qq422016 qtio422016.Writer, name, desc, level string, extra ...string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamcommand(qw422016, name, desc, level, extra...)
	qt422016.ReleaseWriter(qw422016)
}

func command(name, desc, level string, extra ...string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	writecommand(qb422016, name, desc, level, extra...)
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
<span class="tag is-light">Everyone</span>
`)
	case "subs":
		qw422016.N().S(`
<span class="tag is-info">Subs</span>
`)
	case "mods":
		qw422016.N().S(`
<span class="tag is-success">Mods</span>
`)
	case "broadcaster":
		qw422016.N().S(`
<span class="tag is-danger">Broadcaster</span>
`)
	case "admin":
		qw422016.N().S(`
<span class="tag is-black">Admins</span>
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
