package main

import (
	"github.com/pkg/errors"
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

func validateEmail(email string) error {

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
	split := strings.SplitN(email, "@", 1)
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
