// Code generated by goa v3.14.1, DO NOT EDIT.
//
// apisvr HTTP client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package cli

import (
	channelsc "apisvr/services/gen/http/channels/client"
	chatmessagesc "apisvr/services/gen/http/chat_messages/client"
	sessionsc "apisvr/services/gen/http/sessions/client"
	usersc "apisvr/services/gen/http/users/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `channels (list|show|create|update|delete)
chat-messages (list|show|create|update|delete)
sessions (create|delete)
users (list|create)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` channels list --session-id "Inventore expedita dolores quia est."` + "\n" +
		os.Args[0] + ` chat-messages list --limit 2953569855248504038 --channel-id 11643325753069024245 --after 11634181251069557909 --before 6950694595491085628 --session-id "Nesciunt architecto."` + "\n" +
		os.Args[0] + ` sessions create --body '{
      "id_token": "Consectetur suscipit quia est necessitatibus eum quo."
   }'` + "\n" +
		os.Args[0] + ` users list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, any, error) {
	var (
		channelsFlags = flag.NewFlagSet("channels", flag.ContinueOnError)

		channelsListFlags         = flag.NewFlagSet("list", flag.ExitOnError)
		channelsListSessionIDFlag = channelsListFlags.String("session-id", "REQUIRED", "")

		channelsShowFlags         = flag.NewFlagSet("show", flag.ExitOnError)
		channelsShowIDFlag        = channelsShowFlags.String("id", "REQUIRED", "ID")
		channelsShowSessionIDFlag = channelsShowFlags.String("session-id", "REQUIRED", "")

		channelsCreateFlags         = flag.NewFlagSet("create", flag.ExitOnError)
		channelsCreateBodyFlag      = channelsCreateFlags.String("body", "REQUIRED", "")
		channelsCreateSessionIDFlag = channelsCreateFlags.String("session-id", "REQUIRED", "")

		channelsUpdateFlags         = flag.NewFlagSet("update", flag.ExitOnError)
		channelsUpdateBodyFlag      = channelsUpdateFlags.String("body", "REQUIRED", "")
		channelsUpdateIDFlag        = channelsUpdateFlags.String("id", "REQUIRED", "ID")
		channelsUpdateSessionIDFlag = channelsUpdateFlags.String("session-id", "REQUIRED", "")

		channelsDeleteFlags         = flag.NewFlagSet("delete", flag.ExitOnError)
		channelsDeleteIDFlag        = channelsDeleteFlags.String("id", "REQUIRED", "ID")
		channelsDeleteSessionIDFlag = channelsDeleteFlags.String("session-id", "REQUIRED", "")

		chatMessagesFlags = flag.NewFlagSet("chat-messages", flag.ContinueOnError)

		chatMessagesListFlags         = flag.NewFlagSet("list", flag.ExitOnError)
		chatMessagesListLimitFlag     = chatMessagesListFlags.String("limit", "50", "")
		chatMessagesListChannelIDFlag = chatMessagesListFlags.String("channel-id", "", "")
		chatMessagesListAfterFlag     = chatMessagesListFlags.String("after", "", "")
		chatMessagesListBeforeFlag    = chatMessagesListFlags.String("before", "", "")
		chatMessagesListSessionIDFlag = chatMessagesListFlags.String("session-id", "REQUIRED", "")

		chatMessagesShowFlags         = flag.NewFlagSet("show", flag.ExitOnError)
		chatMessagesShowIDFlag        = chatMessagesShowFlags.String("id", "REQUIRED", "ID")
		chatMessagesShowSessionIDFlag = chatMessagesShowFlags.String("session-id", "REQUIRED", "")

		chatMessagesCreateFlags         = flag.NewFlagSet("create", flag.ExitOnError)
		chatMessagesCreateBodyFlag      = chatMessagesCreateFlags.String("body", "REQUIRED", "")
		chatMessagesCreateSessionIDFlag = chatMessagesCreateFlags.String("session-id", "REQUIRED", "")

		chatMessagesUpdateFlags         = flag.NewFlagSet("update", flag.ExitOnError)
		chatMessagesUpdateBodyFlag      = chatMessagesUpdateFlags.String("body", "REQUIRED", "")
		chatMessagesUpdateIDFlag        = chatMessagesUpdateFlags.String("id", "REQUIRED", "ID")
		chatMessagesUpdateSessionIDFlag = chatMessagesUpdateFlags.String("session-id", "REQUIRED", "")

		chatMessagesDeleteFlags         = flag.NewFlagSet("delete", flag.ExitOnError)
		chatMessagesDeleteIDFlag        = chatMessagesDeleteFlags.String("id", "REQUIRED", "ID")
		chatMessagesDeleteSessionIDFlag = chatMessagesDeleteFlags.String("session-id", "REQUIRED", "")

		sessionsFlags = flag.NewFlagSet("sessions", flag.ContinueOnError)

		sessionsCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		sessionsCreateBodyFlag = sessionsCreateFlags.String("body", "REQUIRED", "")

		sessionsDeleteFlags         = flag.NewFlagSet("delete", flag.ExitOnError)
		sessionsDeleteSessionIDFlag = sessionsDeleteFlags.String("session-id", "REQUIRED", "")

		usersFlags = flag.NewFlagSet("users", flag.ContinueOnError)

		usersListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		usersCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		usersCreateBodyFlag = usersCreateFlags.String("body", "REQUIRED", "")
	)
	channelsFlags.Usage = channelsUsage
	channelsListFlags.Usage = channelsListUsage
	channelsShowFlags.Usage = channelsShowUsage
	channelsCreateFlags.Usage = channelsCreateUsage
	channelsUpdateFlags.Usage = channelsUpdateUsage
	channelsDeleteFlags.Usage = channelsDeleteUsage

	chatMessagesFlags.Usage = chatMessagesUsage
	chatMessagesListFlags.Usage = chatMessagesListUsage
	chatMessagesShowFlags.Usage = chatMessagesShowUsage
	chatMessagesCreateFlags.Usage = chatMessagesCreateUsage
	chatMessagesUpdateFlags.Usage = chatMessagesUpdateUsage
	chatMessagesDeleteFlags.Usage = chatMessagesDeleteUsage

	sessionsFlags.Usage = sessionsUsage
	sessionsCreateFlags.Usage = sessionsCreateUsage
	sessionsDeleteFlags.Usage = sessionsDeleteUsage

	usersFlags.Usage = usersUsage
	usersListFlags.Usage = usersListUsage
	usersCreateFlags.Usage = usersCreateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "channels":
			svcf = channelsFlags
		case "chat-messages":
			svcf = chatMessagesFlags
		case "sessions":
			svcf = sessionsFlags
		case "users":
			svcf = usersFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "channels":
			switch epn {
			case "list":
				epf = channelsListFlags

			case "show":
				epf = channelsShowFlags

			case "create":
				epf = channelsCreateFlags

			case "update":
				epf = channelsUpdateFlags

			case "delete":
				epf = channelsDeleteFlags

			}

		case "chat-messages":
			switch epn {
			case "list":
				epf = chatMessagesListFlags

			case "show":
				epf = chatMessagesShowFlags

			case "create":
				epf = chatMessagesCreateFlags

			case "update":
				epf = chatMessagesUpdateFlags

			case "delete":
				epf = chatMessagesDeleteFlags

			}

		case "sessions":
			switch epn {
			case "create":
				epf = sessionsCreateFlags

			case "delete":
				epf = sessionsDeleteFlags

			}

		case "users":
			switch epn {
			case "list":
				epf = usersListFlags

			case "create":
				epf = usersCreateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "channels":
			c := channelsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = channelsc.BuildListPayload(*channelsListSessionIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = channelsc.BuildShowPayload(*channelsShowIDFlag, *channelsShowSessionIDFlag)
			case "create":
				endpoint = c.Create()
				data, err = channelsc.BuildCreatePayload(*channelsCreateBodyFlag, *channelsCreateSessionIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = channelsc.BuildUpdatePayload(*channelsUpdateBodyFlag, *channelsUpdateIDFlag, *channelsUpdateSessionIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = channelsc.BuildDeletePayload(*channelsDeleteIDFlag, *channelsDeleteSessionIDFlag)
			}
		case "chat-messages":
			c := chatmessagesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = chatmessagesc.BuildListPayload(*chatMessagesListLimitFlag, *chatMessagesListChannelIDFlag, *chatMessagesListAfterFlag, *chatMessagesListBeforeFlag, *chatMessagesListSessionIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = chatmessagesc.BuildShowPayload(*chatMessagesShowIDFlag, *chatMessagesShowSessionIDFlag)
			case "create":
				endpoint = c.Create()
				data, err = chatmessagesc.BuildCreatePayload(*chatMessagesCreateBodyFlag, *chatMessagesCreateSessionIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = chatmessagesc.BuildUpdatePayload(*chatMessagesUpdateBodyFlag, *chatMessagesUpdateIDFlag, *chatMessagesUpdateSessionIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = chatmessagesc.BuildDeletePayload(*chatMessagesDeleteIDFlag, *chatMessagesDeleteSessionIDFlag)
			}
		case "sessions":
			c := sessionsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = sessionsc.BuildCreatePayload(*sessionsCreateBodyFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = sessionsc.BuildDeletePayload(*sessionsDeleteSessionIDFlag)
			}
		case "users":
			c := usersc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "create":
				endpoint = c.Create()
				data, err = usersc.BuildCreatePayload(*usersCreateBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// channelsUsage displays the usage of the channels command and its subcommands.
func channelsUsage() {
	fmt.Fprintf(os.Stderr, `Service is the channels service interface.
Usage:
    %[1]s [globalflags] channels COMMAND [flags]

COMMAND:
    list: List implements list.
    show: Show implements show.
    create: Create implements create.
    update: Update implements update.
    delete: Delete implements delete.

Additional help:
    %[1]s channels COMMAND --help
`, os.Args[0])
}
func channelsListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels list -session-id STRING

List implements list.
    -session-id STRING: 

Example:
    %[1]s channels list --session-id "Inventore expedita dolores quia est."
`, os.Args[0])
}

func channelsShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels show -id UINT64 -session-id STRING

Show implements show.
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s channels show --id 928142019106042876 --session-id "Aut soluta molestiae corrupti nihil."
`, os.Args[0])
}

func channelsCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels create -body JSON -session-id STRING

Create implements create.
    -body JSON: 
    -session-id STRING: 

Example:
    %[1]s channels create --body '{
      "name": "Et ipsa."
   }' --session-id "Ullam qui aut veritatis."
`, os.Args[0])
}

func channelsUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels update -body JSON -id UINT64 -session-id STRING

Update implements update.
    -body JSON: 
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s channels update --body '{
      "name": "Ut quibusdam ea quos."
   }' --id 12008964473146260585 --session-id "Ratione et aliquam reiciendis at officiis."
`, os.Args[0])
}

func channelsDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels delete -id UINT64 -session-id STRING

Delete implements delete.
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s channels delete --id 10526203734448781454 --session-id "Deleniti consequatur voluptates ab quia quaerat rerum."
`, os.Args[0])
}

// chat-messagesUsage displays the usage of the chat-messages command and its
// subcommands.
func chatMessagesUsage() {
	fmt.Fprintf(os.Stderr, `Service is the chat_messages service interface.
Usage:
    %[1]s [globalflags] chat-messages COMMAND [flags]

COMMAND:
    list: List implements list.
    show: Show implements show.
    create: Create implements create.
    update: Update implements update.
    delete: Delete implements delete.

Additional help:
    %[1]s chat-messages COMMAND --help
`, os.Args[0])
}
func chatMessagesListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages list -limit INT -channel-id UINT64 -after UINT64 -before UINT64 -session-id STRING

List implements list.
    -limit INT: 
    -channel-id UINT64: 
    -after UINT64: 
    -before UINT64: 
    -session-id STRING: 

Example:
    %[1]s chat-messages list --limit 2953569855248504038 --channel-id 11643325753069024245 --after 11634181251069557909 --before 6950694595491085628 --session-id "Nesciunt architecto."
`, os.Args[0])
}

func chatMessagesShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages show -id UINT64 -session-id STRING

Show implements show.
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s chat-messages show --id 735355215373132696 --session-id "In est similique non impedit."
`, os.Args[0])
}

func chatMessagesCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages create -body JSON -session-id STRING

Create implements create.
    -body JSON: 
    -session-id STRING: 

Example:
    %[1]s chat-messages create --body '{
      "channel_id": 3926127571255298770,
      "content": "In dolore odio quasi quos."
   }' --session-id "Consequatur unde est exercitationem occaecati est fugit."
`, os.Args[0])
}

func chatMessagesUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages update -body JSON -id UINT64 -session-id STRING

Update implements update.
    -body JSON: 
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s chat-messages update --body '{
      "content": "Sit quos enim praesentium provident et."
   }' --id 7805942058009010906 --session-id "Laborum sequi necessitatibus."
`, os.Args[0])
}

func chatMessagesDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages delete -id UINT64 -session-id STRING

Delete implements delete.
    -id UINT64: ID
    -session-id STRING: 

Example:
    %[1]s chat-messages delete --id 178502275960762926 --session-id "Quam velit et."
`, os.Args[0])
}

// sessionsUsage displays the usage of the sessions command and its subcommands.
func sessionsUsage() {
	fmt.Fprintf(os.Stderr, `Service is the sessions service interface.
Usage:
    %[1]s [globalflags] sessions COMMAND [flags]

COMMAND:
    create: Create implements create.
    delete: Delete implements delete.

Additional help:
    %[1]s sessions COMMAND --help
`, os.Args[0])
}
func sessionsCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sessions create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s sessions create --body '{
      "id_token": "Consectetur suscipit quia est necessitatibus eum quo."
   }'
`, os.Args[0])
}

func sessionsDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sessions delete -session-id STRING

Delete implements delete.
    -session-id STRING: 

Example:
    %[1]s sessions delete --session-id "Ipsa et minima quaerat omnis adipisci."
`, os.Args[0])
}

// usersUsage displays the usage of the users command and its subcommands.
func usersUsage() {
	fmt.Fprintf(os.Stderr, `Service is the users service interface.
Usage:
    %[1]s [globalflags] users COMMAND [flags]

COMMAND:
    list: List implements list.
    create: Create implements create.

Additional help:
    %[1]s users COMMAND --help
`, os.Args[0])
}
func usersListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users list

List implements list.

Example:
    %[1]s users list
`, os.Args[0])
}

func usersCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] users create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s users create --body '{
      "email": "Autem eum ducimus beatae.",
      "name": "Sit perferendis temporibus praesentium."
   }'
`, os.Args[0])
}
