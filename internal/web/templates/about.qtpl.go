// Code generated by qtc from "about.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package templates

import "github.com/hortbot/hortbot/internal/version"

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

type AboutPage struct {
	BasePage
}

func (p *AboutPage) StreamPageTitle(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
`)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(` - About
`)
}

func (p *AboutPage) WritePageTitle(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageTitle(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *AboutPage) PageTitle() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageTitle(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}

func (p *AboutPage) StreamPageBody(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<section class="section">
    <div class="container content">
        <h1 class="title has-text-centered">
            About
        </h1>

        <div class="columns">
            <div class="column is-8 is-offset-2 has-text-centered">
                <p>
                    `)
	p.StreamPageBrand(qw422016)
	qw422016.N().S(` is an instance of HortBot, a Twitch chat bot written in Go.
                    You can find its source code on
                    <a href="https://github.com/hortbot/hortbot">GitHub here</a>.
                </p>
                <p>
                    This site is currently running version <code>`)
	qw422016.E().S(version.Version())
	qw422016.N().S(`</code>.
                </p>
            </div>
        </div>
    </div>
</section>
`)
}

func (p *AboutPage) WritePageBody(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	p.StreamPageBody(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func (p *AboutPage) PageBody() string {
	qb422016 := qt422016.AcquireByteBuffer()
	p.WritePageBody(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
