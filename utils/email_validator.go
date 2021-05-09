package utils

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
	"net"
	"strings"
)

//<?php
//
////http://www.linuxjournal.com/article/9585
///**
//Validate an email address.
//Provide email address (raw input)
//Returns true if the email address has the email
//address format and the domain exists.
//*/

var emailCache = make(map[string]bool)

func GoogleDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "8.8.8.8:53")
}

func CloudflareDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "1.1.1.1:53")
}

func ValidateEmail(email string) (err error) {

	if email == "" {
		return errors.New("email address is empty")
	}

	//function validateEmail($email)
	//{
	//$isValid = true;
	//$atIndex = strrpos($email, "@");
	//if (is_bool($atIndex) && !$atIndex)
	//{
	//$isValid = false;
	//}
	if !strings.Contains(email, "@") {
		return errors.New("emails doesn't contain an @")
	}

	//else
	//{
	//$domain = substr($email, $atIndex+1);
	//$local = substr($email, 0, $atIndex);
	//$localLen = strlen($local);
	//$domainLen = strlen($domain);
	//if ($localLen < 1 || $localLen > 64)
	//{
	//// local part length exceeded
	//$isValid = false;
	//}
	split := strings.SplitN(email, "@", 2)
	domain := split[1]
	local := split[0]
	domainLen := len(domain)
	localLen := len(local)

	if localLen < 1 {
		return errors.New("email no user part")
	}
	if localLen > 64 {
		return errors.New("email user part exceeds max length")
	}

	//else if ($domainLen < 1 || $domainLen > 255)
	//{
	//// domain part length exceeded
	//$isValid = false;
	//}

	if domainLen < 1 {
		return errors.New("email no domain part")
	}
	if domainLen > 255 {
		return errors.New("email domain part exceeds max length")
	}

	// Try to respond from the cache first
	value, ok := emailCache[domain]
	if ok && value == true {
		return nil
	}
	if ok && value == false {
		return errors.New("email previously declined")
	}
	//defer func() {
	//	if err != nil {
	//		emailCache[domain] = false
	//	}
	//}()

	// 1. try local dns
	mxs, err := net.LookupMX(domain)
	if err != nil {

		r := net.Resolver{
			PreferGo: true,
			Dial:     CloudflareDNSDialer,
		}
		ctx := context.Background()
		ips, err := r.LookupIPAddr(ctx, domain)
		if err != nil {
			r := net.Resolver{
				PreferGo: true,
				Dial:     GoogleDNSDialer,
			}
			ctx := context.Background()
			ips, err := r.LookupIPAddr(ctx, domain)
			if err != nil {
				emailCache[domain] = false
				return errors.New("could not find email domain")
			}

			for _, ip := range ips {
				log.Printf("(Google) mail server: %s", ip)
			}
		}

		for _, ip := range ips {
			log.Printf("(Cloudflare) mail server: %s", ip)
		}
	}

	for _, mx := range mxs {
		log.Printf("(Local) mail server: %s", mx.Host)
	}

	emailCache[domain] = true
	return nil
}

//else if ($local[0] == '.' || $local[$localLen-1] == '.')
//{
//// local part starts or ends with '.'
//$isValid = false;
//}
//else if (preg_match('/\\.\\./', $local))
//{
//// local part has two consecutive dots
//$isValid = false;
//}
//else if (!preg_match('/^[A-Za-z0-9\\-\\.]+$/', $domain))
//{
//// character not valid in domain part
//$isValid = false;
//}
//else if (preg_match('/\\.\\./', $domain))
//{
//// domain part has two consecutive dots
//$isValid = false;
//}
//else if
//(!preg_match('/^(\\\\.|[A-Za-z0-9!#%&`_=\\/$\'*+?^{}|~.-])+$/',
//                 str_replace("\\\\","",$local)))
//      {
//         // character not valid in local part unless
//         // local part is quoted
//         if (!preg_match('/^"(\\\\"|[^"])+"$/',
//             str_replace("\\\\","",$local)))
//         {
//            $isValid = false;
//         }
//      }
//      if ($isValid && !(checkdnsrr($domain,"MX") || checkdnsrr($domain,"A")))
//      {
//         // domain not found in DNS
//         $isValid = false;
//      }
//   }
//   return $isValid;
//}
//
//?>
