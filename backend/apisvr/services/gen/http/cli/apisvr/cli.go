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
	notificationsc "apisvr/services/gen/http/notifications/client"
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
notifications subscribe
users (list|create)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` channels list --id-token "abcdef12345"` + "\n" +
		os.Args[0] + ` chat-messages list --limit 4957638961049071245 --channel-id 248369442570821029 --after 11502104032657031607 --before 821952026860172930 --id-token "abcdef12345"` + "\n" +
		os.Args[0] + ` notifications subscribe --id-token "abcdef12345"` + "\n" +
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
	dialer goahttp.Dialer,
	notificationsConfigurer *notificationsc.ConnConfigurer,
) (goa.Endpoint, any, error) {
	var (
		channelsFlags = flag.NewFlagSet("channels", flag.ContinueOnError)

		channelsListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		channelsListIDTokenFlag = channelsListFlags.String("id-token", "REQUIRED", "")

		channelsShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		channelsShowIDFlag      = channelsShowFlags.String("id", "REQUIRED", "ID")
		channelsShowIDTokenFlag = channelsShowFlags.String("id-token", "REQUIRED", "")

		channelsCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		channelsCreateBodyFlag    = channelsCreateFlags.String("body", "REQUIRED", "")
		channelsCreateIDTokenFlag = channelsCreateFlags.String("id-token", "REQUIRED", "")

		channelsUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		channelsUpdateBodyFlag    = channelsUpdateFlags.String("body", "REQUIRED", "")
		channelsUpdateIDFlag      = channelsUpdateFlags.String("id", "REQUIRED", "ID")
		channelsUpdateIDTokenFlag = channelsUpdateFlags.String("id-token", "REQUIRED", "")

		channelsDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		channelsDeleteIDFlag      = channelsDeleteFlags.String("id", "REQUIRED", "ID")
		channelsDeleteIDTokenFlag = channelsDeleteFlags.String("id-token", "REQUIRED", "")

		chatMessagesFlags = flag.NewFlagSet("chat-messages", flag.ContinueOnError)

		chatMessagesListFlags         = flag.NewFlagSet("list", flag.ExitOnError)
		chatMessagesListLimitFlag     = chatMessagesListFlags.String("limit", "50", "")
		chatMessagesListChannelIDFlag = chatMessagesListFlags.String("channel-id", "", "")
		chatMessagesListAfterFlag     = chatMessagesListFlags.String("after", "", "")
		chatMessagesListBeforeFlag    = chatMessagesListFlags.String("before", "", "")
		chatMessagesListIDTokenFlag   = chatMessagesListFlags.String("id-token", "REQUIRED", "")

		chatMessagesShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		chatMessagesShowIDFlag      = chatMessagesShowFlags.String("id", "REQUIRED", "ID")
		chatMessagesShowIDTokenFlag = chatMessagesShowFlags.String("id-token", "REQUIRED", "")

		chatMessagesCreateFlags       = flag.NewFlagSet("create", flag.ExitOnError)
		chatMessagesCreateBodyFlag    = chatMessagesCreateFlags.String("body", "REQUIRED", "")
		chatMessagesCreateIDTokenFlag = chatMessagesCreateFlags.String("id-token", "REQUIRED", "")

		chatMessagesUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		chatMessagesUpdateBodyFlag    = chatMessagesUpdateFlags.String("body", "REQUIRED", "")
		chatMessagesUpdateIDFlag      = chatMessagesUpdateFlags.String("id", "REQUIRED", "ID")
		chatMessagesUpdateIDTokenFlag = chatMessagesUpdateFlags.String("id-token", "REQUIRED", "")

		chatMessagesDeleteFlags       = flag.NewFlagSet("delete", flag.ExitOnError)
		chatMessagesDeleteIDFlag      = chatMessagesDeleteFlags.String("id", "REQUIRED", "ID")
		chatMessagesDeleteIDTokenFlag = chatMessagesDeleteFlags.String("id-token", "REQUIRED", "")

		notificationsFlags = flag.NewFlagSet("notifications", flag.ContinueOnError)

		notificationsSubscribeFlags       = flag.NewFlagSet("subscribe", flag.ExitOnError)
		notificationsSubscribeIDTokenFlag = notificationsSubscribeFlags.String("id-token", "REQUIRED", "")

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

	notificationsFlags.Usage = notificationsUsage
	notificationsSubscribeFlags.Usage = notificationsSubscribeUsage

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
		case "notifications":
			svcf = notificationsFlags
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

		case "notifications":
			switch epn {
			case "subscribe":
				epf = notificationsSubscribeFlags

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
				data, err = channelsc.BuildListPayload(*channelsListIDTokenFlag)
			case "show":
				endpoint = c.Show()
				data, err = channelsc.BuildShowPayload(*channelsShowIDFlag, *channelsShowIDTokenFlag)
			case "create":
				endpoint = c.Create()
				data, err = channelsc.BuildCreatePayload(*channelsCreateBodyFlag, *channelsCreateIDTokenFlag)
			case "update":
				endpoint = c.Update()
				data, err = channelsc.BuildUpdatePayload(*channelsUpdateBodyFlag, *channelsUpdateIDFlag, *channelsUpdateIDTokenFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = channelsc.BuildDeletePayload(*channelsDeleteIDFlag, *channelsDeleteIDTokenFlag)
			}
		case "chat-messages":
			c := chatmessagesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = chatmessagesc.BuildListPayload(*chatMessagesListLimitFlag, *chatMessagesListChannelIDFlag, *chatMessagesListAfterFlag, *chatMessagesListBeforeFlag, *chatMessagesListIDTokenFlag)
			case "show":
				endpoint = c.Show()
				data, err = chatmessagesc.BuildShowPayload(*chatMessagesShowIDFlag, *chatMessagesShowIDTokenFlag)
			case "create":
				endpoint = c.Create()
				data, err = chatmessagesc.BuildCreatePayload(*chatMessagesCreateBodyFlag, *chatMessagesCreateIDTokenFlag)
			case "update":
				endpoint = c.Update()
				data, err = chatmessagesc.BuildUpdatePayload(*chatMessagesUpdateBodyFlag, *chatMessagesUpdateIDFlag, *chatMessagesUpdateIDTokenFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = chatmessagesc.BuildDeletePayload(*chatMessagesDeleteIDFlag, *chatMessagesDeleteIDTokenFlag)
			}
		case "notifications":
			c := notificationsc.NewClient(scheme, host, doer, enc, dec, restore, dialer, notificationsConfigurer)
			switch epn {
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = notificationsc.BuildSubscribePayload(*notificationsSubscribeIDTokenFlag)
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels list -id-token STRING

List implements list.
    -id-token STRING: 

Example:
    %[1]s channels list --id-token "abcdef12345"
`, os.Args[0])
}

func channelsShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels show -id UINT64 -id-token STRING

Show implements show.
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s channels show --id 13449821290766471387 --id-token "abcdef12345"
`, os.Args[0])
}

func channelsCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels create -body JSON -id-token STRING

Create implements create.
    -body JSON: 
    -id-token STRING: 

Example:
    %[1]s channels create --body '{
      "name": "Dolorem optio consequatur est eligendi."
   }' --id-token "abcdef12345"
`, os.Args[0])
}

func channelsUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels update -body JSON -id UINT64 -id-token STRING

Update implements update.
    -body JSON: 
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s channels update --body '{
      "name": "Animi ut aut totam."
   }' --id 5188841799517865761 --id-token "abcdef12345"
`, os.Args[0])
}

func channelsDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] channels delete -id UINT64 -id-token STRING

Delete implements delete.
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s channels delete --id 7140320228966239079 --id-token "abcdef12345"
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
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages list -limit INT -channel-id UINT64 -after UINT64 -before UINT64 -id-token STRING

List implements list.
    -limit INT: 
    -channel-id UINT64: 
    -after UINT64: 
    -before UINT64: 
    -id-token STRING: 

Example:
    %[1]s chat-messages list --limit 4957638961049071245 --channel-id 248369442570821029 --after 11502104032657031607 --before 821952026860172930 --id-token "abcdef12345"
`, os.Args[0])
}

func chatMessagesShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages show -id UINT64 -id-token STRING

Show implements show.
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s chat-messages show --id 9267441921255015593 --id-token "abcdef12345"
`, os.Args[0])
}

func chatMessagesCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages create -body JSON -id-token STRING

Create implements create.
    -body JSON: 
    -id-token STRING: 

Example:
    %[1]s chat-messages create --body '{
      "channel_id": 3412152620045203249,
      "content": "Non accusantium."
   }' --id-token "abcdef12345"
`, os.Args[0])
}

func chatMessagesUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages update -body JSON -id UINT64 -id-token STRING

Update implements update.
    -body JSON: 
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s chat-messages update --body '{
      "content": "Qui similique dolor."
   }' --id 9594253205796740350 --id-token "abcdef12345"
`, os.Args[0])
}

func chatMessagesDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chat-messages delete -id UINT64 -id-token STRING

Delete implements delete.
    -id UINT64: ID
    -id-token STRING: 

Example:
    %[1]s chat-messages delete --id 14096741058994480104 --id-token "abcdef12345"
`, os.Args[0])
}

// notificationsUsage displays the usage of the notifications command and its
// subcommands.
func notificationsUsage() {
	fmt.Fprintf(os.Stderr, `Service is the notifications service interface.
Usage:
    %[1]s [globalflags] notifications COMMAND [flags]

COMMAND:
    subscribe: Subscribe to events sent such new chat messages.

Additional help:
    %[1]s notifications COMMAND --help
`, os.Args[0])
}
func notificationsSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] notifications subscribe -id-token STRING

Subscribe to events sent such new chat messages.
    -id-token STRING: 

Example:
    %[1]s notifications subscribe --id-token "abcdef12345"
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
      "email": "Vero in dolore odio quasi.",
      "name": "Quo minus quas voluptatibus consequatur nemo."
   }'
`, os.Args[0])
}
