// Code generated by templ - DO NOT EDIT.

package productbacklog

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "tasklify/internal/database"
import "fmt"

func productBacklog(backlogUserStories []database.UserStory, sprints map[string]database.Sprint, projectID uint) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"dialog\"><div class=\"bg-base-100 min-h-screen\"><div><!-- Buttons Container --><div class=\"flex justify-between items-center mb-4 pt-2 pl-2 pr-2\"><!-- Left Aligned Buttons --><div><a href=\"/dashboard\" class=\"btn btn-primary btn-sm\">Back to Dashboard</a></div><!-- Right Aligned Buttons --><div><form hx-get=\"/createuserstory\" hx-target=\"#dialog-create\"><input type=\"hidden\" id=\"projectID\" name=\"projectID\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprint(projectID)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <button class=\"btn btn-primary btn-sm\">Create User Story</button></form><button class=\"btn btn-primary btn-sm ml-2\" hx-get=\"/createsprint\" hx-target=\"#dialog-create\">Create Sprint</button></div></div><div id=\"dialog-create\"></div><div class=\"flex flex-col w-full lg:flex-row\"><!-- Backlog - unassigned, unrealized --><div class=\"card card-compact bg-base-200 transition-shadow w-2/5 mt-2 ml-2\"><div class=\"card-title collapse-title text-xl font-medium bg-base-300 border-base-300\">Product backlog</div><div class=\"card-body\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, us := range backlogUserStories {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"card card-compact bg-base-100 shadow-xl hover:shadow-2xl transition-shadow w-full\"><div class=\"card-body\"><h2 class=\"card-title\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(us.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 33, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(*us.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 34, Col: 30}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Divide --><div class=\"divider lg:divider-horizontal\"></div><!-- Sprints --><div class=\"bg-base-100 join join-vertical w-full mt-2 mr-2\"><!-- Sprint -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, sprint := range sprints {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"collapse collapse-arrow border border-base-300 mb-4\"><input type=\"checkbox\" checked><div class=\"collapse-title text-xl font-medium bg-base-300 text-base-500\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(sprint.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 49, Col: 23}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if sprint.DetermineStatus() == database.StatusInProgress {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-red-300\">Active</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if sprint.DetermineStatus() == database.StatusDone {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-300\">Done</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"ml-2 inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-300\">Done</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"collapse-content card-body bg-base-100\"><div class=\"flex justify-end\"><!-- Places the button nicely --><!-- Sprint Backlog Button -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if sprint.DetermineStatus() == database.StatusInProgress {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-get=\"/sprintbacklog\" hx-target=\"#whole\" class=\"ml-auto\"><button class=\"btn btn-sm bg-primary\">Sprint Backlog</button> <input type=\"hidden\" id=\"sprintID\" name=\"sprintID\" value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprint(sprint.ID)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></form>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- user story table and other contents --><div class=\"overflow-x-auto\"><table class=\"table\"><!-- head --><thead><tr><th>Title</th><th>Description</th><th>Priority</th><th>Assigned</th><th>Realized</th><th></th></tr></thead> <tbody><!-- User stories rows -->")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, v := range sprint.UserStories {
				templ_7745c5c3_Err = userStoryTableRow(v).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</tbody></table></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func userStoryCard(name string, description string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- user story --><div class=\"card card-compact bg-base-100 shadow-xl hover:shadow-2xl transition-shadow w-full\"><div class=\"card-body\"><h2 class=\"card-title\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 111, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 112, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func userStoryTableRow(us database.UserStory) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- row 1 --><tr class=\"hover\"><td><div class=\"font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(us.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 121, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></td><!-- Description --><td class=\"min-w-[12rem] max-w-[20rem] whitespace-normal\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(*us.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/web/pages/productbacklog/productBacklog.templ`, Line: 125, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><!-- Priority --><td>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch us.Priority {
		case database.PriorityMustHave:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"badge badge-error\">Must have</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case database.PriorityCouldHave:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"badge badge-warning\">Could have</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case database.PriorityShouldHave:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"badge badge-success\">Should have</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case database.PriorityWontHaveThisTime:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"badge badge-info\">Won't have this time</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		default:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span></span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><!-- Assigned --><td>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if us.UserID == nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18 18 6M6 6l12 12\"></path></svg>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m4.5 12.75 6 6 9-13.5\"></path></svg>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><!-- Realized --><td>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if *us.Realized {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m4.5 12.75 6 6 9-13.5\"></path></svg>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"w-6 h-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18 18 6M6 6l12 12\"></path></svg>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><th><button class=\"btn btn-xs\">details</button></th></tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
