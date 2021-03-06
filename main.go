package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
		r := gin.Default()
		r.GET("/ticket", otrs)
		r.Run()
}

func otrs(c *gin.Context) {
	var ticket TicketResponse

	ticketId := c.Query("ticket_id")
	login    := c.Query("login")
	pass     := c.Query("pass")
	url 	 := c.Query("url")

	resp, err := http.Get("http://"+url+"/otrs/nph-genericinterface.pl/Webservice/GenericTicketConnectorREST/TicketList?UserLogin="+login+"&Password="+pass+"&TicketID="+ticketId)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &ticket)

	c.JSON(200, gin.H{
		"message": ticket,
	})
}

type Ticket struct {
	ResponsibleID  string `json:"ResponsibleID"`
	Owner          string `json:"Owner"`
	Created        string `json:"Created"`
	StateType      string `json:"StateType"`
	Type           string `json:"Type"`
	CreateTimeUnix string `json:"CreateTimeUnix"`
	State          string `json:"State"`
	OwnerID        string `json:"OwnerID"`
	UnlockTimeout  string `json:"UnlockTimeout"`
	Priority       string `json:"Priority"`
	GroupID        string `json:"GroupID"`
	StateID        string `json:"StateID"`
	SLAID          string `json:"SLAID"`
	CustomerUserID string `json:"CustomerUserID"`
	Article        []struct {
		ResponsibleID          string      `json:"ResponsibleID"`
		Owner                  string      `json:"Owner"`
		InReplyTo              string      `json:"InReplyTo"`
		From                   string      `json:"From"`
		ContentCharset         string      `json:"ContentCharset"`
		Created                string      `json:"Created"`
		SenderTypeID           string      `json:"SenderTypeID"`
		AgeTimeUnix            int         `json:"AgeTimeUnix"`
		To                     string      `json:"To"`
		References             string      `json:"References"`
		StateType              string      `json:"StateType"`
		Type                   string      `json:"Type"`
		ContentType            string      `json:"ContentType"`
		CreateTimeUnix         string      `json:"CreateTimeUnix"`
		State                  string      `json:"State"`
		OwnerID                string      `json:"OwnerID"`
		Priority               string      `json:"Priority"`
		FromRealname           string      `json:"FromRealname"`
		StateID                string      `json:"StateID"`
		SLAID                  interface{} `json:"SLAID"`
		Subject                string      `json:"Subject"`
		CreatedBy              string      `json:"CreatedBy"`
		CustomerUserID         string      `json:"CustomerUserID"`
		ReplyTo                string      `json:"ReplyTo"`
		Title                  string      `json:"Title"`
		ServiceID              interface{} `json:"ServiceID"`
		Body                   string      `json:"Body"`
		PriorityID             string      `json:"PriorityID"`
		SenderType             string      `json:"SenderType"`
		SLA                    string      `json:"SLA"`
		EscalationSolutionTime string      `json:"EscalationSolutionTime"`
		IncomingTime           string      `json:"IncomingTime"`
		Queue                  string      `json:"Queue"`
		CustomerID             string      `json:"CustomerID"`
		QueueID                string      `json:"QueueID"`
		ArticleTypeID          string      `json:"ArticleTypeID"`
		TypeID                 string      `json:"TypeID"`
		ArticleType            string      `json:"ArticleType"`
		Age                    int         `json:"Age"`
		EscalationUpdateTime   string      `json:"EscalationUpdateTime"`
		Charset                string      `json:"Charset"`
		Cc                     string      `json:"Cc"`
		UntilTime              int         `json:"UntilTime"`
		EscalationTime         string      `json:"EscalationTime"`
		TicketID               string      `json:"ticketID"`
		Lock                   string      `json:"Lock"`
		MimeType               string      `json:"MimeType"`
		TicketNumber           string      `json:"TicketNumber"`
		Service                string      `json:"Service"`
		EscalationResponseTime string      `json:"EscalationResponseTime"`
		ToRealname             string      `json:"ToRealname"`
		Responsible            string      `json:"Responsible"`
		RealTillTimeNotUsed    string      `json:"RealTillTimeNotUsed"`
		ArticleID              string      `json:"ArticleID"`
		Changed                string      `json:"Changed"`
		MessageID              string      `json:"MessageID"`
		LockID                 string      `json:"LockID"`
	} `json:"Article"`
	Title                  string `json:"Title"`
	CreateBy               string `json:"CreateBy"`
	ServiceID              string `json:"ServiceID"`
	PriorityID             string `json:"PriorityID"`
	EscalationSolutionTime string `json:"EscalationSolutionTime"`
	Queue                  string `json:"Queue"`
	CustomerID             string `json:"CustomerID"`
	QueueID                string `json:"QueueID"`
	TypeID                 string `json:"TypeID"`
	ArchiveFlag            string `json:"ArchiveFlag"`
	EscalationUpdateTime   string `json:"EscalationUpdateTime"`
	Age                    int    `json:"Age"`
	UntilTime              int    `json:"UntilTime"`
	EscalationTime         string `json:"EscalationTime"`
	TicketID               string `json:"ticketID"`
	Lock                   string `json:"Lock"`
	TicketNumber           string `json:"TicketNumber"`
	EscalationResponseTime string `json:"EscalationResponseTime"`
	ChangeBy               string `json:"ChangeBy"`
	RealTillTimeNotUsed    string `json:"RealTillTimeNotUsed"`
	Responsible            string `json:"Responsible"`
	Changed                string `json:"Changed"`
	LockID                 string `json:"LockID"`
}

type TicketResponse struct {
	Ticket []Ticket `json:"Ticket"`
}