package api

import (
	"testing"
)

func TestValidateWhyString(t *testing.T) {
	type args struct {
		attribute string
		why       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				"why",
				"",
			},
			wantErr: true,
		},
		{
			name: "valid string",
			args: args{
				"why",
				"asdas1212",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateWhyString(tt.args.attribute, tt.args.why); (err != nil) != tt.wantErr {
				t.Errorf("ValidateWhyString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple tag a=a",
			args{
				"a=a",
			},
			false,
		},
		{
			"value is string",
			args{
				"name=value",
			},
			false,
		},
		{
			"value is number",
			args{
				"name=123.12",
			},
			false,
		},
		{
			"value is has space",
			args{
				"name=test 1",
			},
			false,
		},
		{
			"tag value is utf8 character",
			args{
				"name=utf8-_!@#%&\" (*)+.,/$!:;<>=?{}~",
			},
			false,
		},
		{
			"key contains hyphen and underscore",
			args{
				"internal_name:demo-1234=aporeliable_arvind_@",
			},
			false,
		},
		{
			"tag contains 2 equals",
			args{
				"a=a=b",
			},
			false,
		},
		{
			"tag contains 2 touching equals",
			args{
				"a==b",
			},
			false,
		},
		{
			"complex key",
			args{
				`utf8-_!@#%&"(*)+.,/$!:;<>=?{}~=utf8-_!@#%&" (*)+.,/$!:;<>=?{}~`,
			},
			false,
		},

		// Error
		{
			"just a word",
			args{
				"justaword",
			},
			true,
		},
		{
			"key contains spaces",
			args{
				"the key=value",
			},
			true,
		},
		{
			"key starts with =",
			args{
				"=thekey=value",
			},
			true,
		},

		// tag size
		{
			"tags is exactly 1024 bytes",
			args{
				"l=ah6ooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exactly 1026 bytes",
			args{
				"l=ah6ooPaelaizohakee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exacty 1023 bytes",
			args{
				"l=ahooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTag("tag", tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTags(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple tag a=a",
			args{
				[]string{"a=a", "b=b"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTags("tag", tt.args.tags); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTagsWithoutReservedPrefixes(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// success case
		{
			"success case a=a, b=b",
			args{
				[]string{"a=a", "b=b"},
			},
			false,
		},
		// error case
		{
			"simple tag @a=a, b=b",
			args{
				[]string{"@a=a", "b=b"},
			},
			true,
		},
		// error case
		{
			"simple tag +a=a, b=b",
			args{
				[]string{"@a=a", "b=b"},
			},
			true,
		},
		// error case
		{
			"simple tag b=b, $a=a",
			args{
				[]string{"b=b", "@a=a"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTagsWithoutReservedPrefixes("tag", tt.args.tags); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateFQDNList(t *testing.T) {
	type args struct {
		attribute string
		FQDNs     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"www.google.com", "mail.google.com", "host"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFQDNList(tt.args.attribute, tt.args.FQDNs); (err != nil) != tt.wantErr {
				t.Errorf("ValidateFQDNList() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalFQDNList(t *testing.T) {
	type args struct {
		attribute string
		FQDNs     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"www.google.com", "mail.google.com"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalFQDNList(tt.args.attribute, tt.args.FQDNs); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalFQDNList() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDRList(t *testing.T) {
	type args struct {
		attribute string
		cidr      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				[]string{"10.0.0.0/8", "172.16.0.0/12"},
			},
			false,
		},
		// invalid empty
		{
			"invalid empty",
			args{
				"cidr",
				[]string{},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDRList(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateCIDRList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalCIDR(t *testing.T) {
	type args struct {
		attribute string
		cidr      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				[]string{"10.0.0.0/8"},
			},
			false,
		},

		// valid empty
		{
			"empty CIDR list",
			args{
				"cidr",
				[]string{},
			},
			false,
		},

		// invalid
		{
			"invalid CIDR",
			args{
				"cidr",
				[]string{"foo"},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalCIDRList(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateOptionalCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalURLs(t *testing.T) {
	type args struct {
		attribute string
		cidr      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid URLs",
			args{
				"urls",
				[]string{"http://google.com/", "https://www.google.com"},
			},
			false,
		},

		// valid empty
		{
			"empty URL list",
			args{
				"urls",
				[]string{},
			},
			false,
		},

		// invalid
		{
			"invalid URLs",
			args{
				"urls",
				[]string{"/foo.txt"},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalURLs(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalURLs() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidateURLs(t *testing.T) {
	type args struct {
		attribute string
		cidr      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid URLs",
			args{
				"urls",
				[]string{"http://www.google.com/", "https://www.google.com"},
			},
			false,
		},

		// invalid empty
		{
			"empty URL list",
			args{
				"urls",
				[]string{},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateURLs(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("ValidateURLs() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalProtoPorts(t *testing.T) {
	type args struct {
		attribute  string
		protoports []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid protoports",
			args{
				"protoports",
				[]string{"TCP:80,443", "udp:53"},
			},
			false,
		},
		// valid empty
		{
			"empty protoports",
			args{
				"protoports",
				[]string{},
			},
			false,
		},
		// invalid
		{
			"no ports",
			args{
				"protoports",
				[]string{"tcp:"},
			},
			true,
		},
		// invalid
		{
			"invalid port",
			args{
				"protoports",
				[]string{"tcp:foo"},
			},
			true,
		},
		// invalid
		{
			"negative value",
			args{
				"protoports",
				[]string{"tcp:-5"},
			},
			true,
		},
		// invalid
		{
			"max port value",
			args{
				"protoports",
				[]string{"tcp:65536"},
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalProtoPorts(tt.args.attribute, tt.args.protoports); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalProtoPorts() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidatePort(t *testing.T) {
	type args struct {
		attribute string
		port      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid min value
		{
			"valid port min",
			args{
				"port",
				0,
			},
			false,
		},
		// valid max value
		{
			"valid port max",
			args{
				"port",
				65535,
			},
			false,
		},
		// invalid less than min
		{
			"smaller than min value",
			args{
				"port",
				-1,
			},
			true,
		},
		// invalid more than max
		{
			"greater than max value",
			args{
				"port",
				65536,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePort(tt.args.attribute, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePort() test = %s error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidateAvailabilityZone(t *testing.T) {
	type args struct {
		attribute        string
		availabilityzone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				"availabilityzone",
				"",
			},
			wantErr: true,
		},
		{
			name: "valid string",
			args: args{
				"availabilityzone",
				"us-east-1a",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAvailabilityZone(tt.args.attribute, tt.args.availabilityzone); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAvailabilityZone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAwsNetworkServices(t *testing.T) {
	type args struct {
		attribute       string
		networkServices []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty list valid",
			args: args{
				"networkServices",
				[]string{},
			},
			wantErr: false,
		},
		{
			name: "valid list",
			args: args{
				"networkServices",
				[]string{"amazon-dns"},
			},
			wantErr: false,
		},
		{
			name: "invalid due to invalid item in list",
			args: args{
				"networkServices",
				[]string{"amazon-nat"},
			},
			wantErr: true,
		},
		{
			name: "invalid due to mixing invalid item with valid item in list",
			args: args{
				"networkServices",
				[]string{"amazon-nat", "amazon-dns"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAwsNetworkServices(tt.args.attribute, tt.args.networkServices); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAwsNetworkServices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateVPCID(t *testing.T) {
	type args struct {
		attribute string
		vpcid     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				"vpcid",
				"",
			},
			wantErr: true,
		},
		{
			name: "invalid string that does not have 'vpc-'",
			args: args{
				"vpcid",
				"vpc12345",
			},
			wantErr: true,
		},
		{
			name: "invalid string that has 'vpc-' but not in beginning",
			args: args{
				"vpcid",
				"vpc1vpc-12345",
			},
			wantErr: true,
		},
		{
			name: "valid string that starts with 'vpc-'",
			args: args{
				"vpcid",
				"vpc-12345",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateVPCID(tt.args.attribute, tt.args.vpcid); (err != nil) != tt.wantErr {
				t.Errorf("ValidateVPCID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEnis(t *testing.T) {
	type args struct {
		attribute string
		enis      []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty list is invalid",
			args: args{
				"enis",
				[]string{},
			},
			wantErr: true,
		},
		{
			name: "invalid due to invalid item in list",
			args: args{
				"enis",
				[]string{"myeni"},
			},
			wantErr: true,
		},
		{
			name: "invalid due to mixing invalid item with valid items in list",
			args: args{
				"networkServices",
				[]string{"myeni", "eni-1", "eni-2"},
			},
			wantErr: true,
		},
		{
			name: "valid list that has all entries starting with 'eni-'",
			args: args{
				"enis",
				[]string{"eni-1", "eni-2", "eni-3"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEnis(tt.args.attribute, tt.args.enis); (err != nil) != tt.wantErr {
				t.Errorf("ValidateEnis() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateVpcInfo(t *testing.T) {
	type args struct {
		attribute                  string
		VPCAvailabilityZoneSubnets []*VPCAvailabilityZoneSubnet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty list is invalid",
			args: args{
				"VPCAvailabilityZoneSubnet",
				[]*VPCAvailabilityZoneSubnet{},
			},
			wantErr: true,
		},
		{
			name: "empty embedded list is invalid",
			args: args{
				"VPCAvailabilityZoneSubnet",
				[]*VPCAvailabilityZoneSubnet{
					{
						VPCID:            "vpc1",
						SubnetInterfaces: []*AvailabilityZoneSubnetInterface{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid list",
			args: args{
				"VPCAvailabilityZoneSubnet",
				[]*VPCAvailabilityZoneSubnet{
					{
						VPCID: "vpc1",
						SubnetInterfaces: []*AvailabilityZoneSubnetInterface{
							{
								AvailabilityZone:        "us-west-1a",
								SourceNetworkInterfaces: []string{},
								SubnetCIDR:              "10.0.2.0/24",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateVpcInfo(tt.args.attribute, tt.args.VPCAvailabilityZoneSubnets); (err != nil) != tt.wantErr {
				t.Errorf("ValidateVpcInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateVpcSubnetInfo(t *testing.T) {
	type args struct {
		attribute      string
		VpcUsedSubnets []*VpcUsedSubnet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty VpcUsedSubnets is invalid",
			args: args{
				"VpcUsedSubnets",
				[]*VpcUsedSubnet{},
			},
			wantErr: true,
		},
		{
			name: "empty AvailabilityZones is invalid",
			args: args{
				"VpcUsedSubnets",
				[]*VpcUsedSubnet{
					{
						VPCID:             "vpc1",
						AvailabilityZones: []string{},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "valid list",
			args: args{
				"VpcUsedSubnets",
				[]*VpcUsedSubnet{
					{
						VPCID:             "vpc1",
						AvailabilityZones: []string{"us-west-1a"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateVpcSubnetInfo(tt.args.attribute, tt.args.VpcUsedSubnets); (err != nil) != tt.wantErr {
				t.Errorf("ValidateVpcSubnetInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
