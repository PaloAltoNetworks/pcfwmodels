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

func TestValidateMirrorRules(t *testing.T) {
	type args struct {
		attribute   string
		mirrorrules []*MirrorRule
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			// invalid ports since protocol is not tcp or udp
			"invalid protocol for ports",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 0,
						DestinationToPort:   650,
						Protocol:            10,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"valid protocol but invalid ports",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 0,
						DestinationToPort:   65536,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"invalid destination port range due to start being -1",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: -1,
						DestinationToPort:   65536,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"invalid destination port range due to end being -1",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 1,
						DestinationToPort:   -1,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"invalid source port range due to start being -1",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 1,
						DestinationToPort:   65536,
						Protocol:            17,
						SourceFromPort:      -1,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"invalid source port range due to end being -1",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 1,
						DestinationToPort:   10,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        -1,
					},
				},
			},
			true,
		},
		{
			"valid protocol but destination port start is greater than end",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 1000,
						DestinationToPort:   10,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"valid protocol but source port start is greater than end",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 10,
						DestinationToPort:   1000,
						Protocol:            6,
						SourceFromPort:      110,
						SourceToPort:        10,
					},
				},
			},
			true,
		},
		{
			"one valid and one invalid rule",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 10,
						DestinationToPort:   165,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
					{
						DestinationFromPort: 100,
						DestinationToPort:   0,
						Protocol:            6,
						SourceFromPort:      10,
						SourceToPort:        100,
					},
				},
			},
			true,
		},
		{
			"valid protocols and valid ports",
			args{
				"mirrorrules",
				[]*MirrorRule{
					{
						DestinationFromPort: 10,
						DestinationToPort:   165,
						Protocol:            17,
						SourceFromPort:      0,
						SourceToPort:        10,
					},
					{
						DestinationFromPort: 0,
						DestinationToPort:   0,
						Protocol:            6,
						SourceFromPort:      10,
						SourceToPort:        100,
					},
					{
						DestinationFromPort: -1,
						DestinationToPort:   -1,
						Protocol:            100,
						SourceFromPort:      -1,
						SourceToPort:        -1,
					},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateMirrorRules(tt.args.attribute, tt.args.mirrorrules); (err != nil) != tt.wantErr {
				t.Errorf("ValidateMirrorRules() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAthenaWorkGroup(t *testing.T) {
	type args struct {
		attribute       string
		athenaWorkGroup string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty AthenaWorkGroup is invalid",
			args: args{
				"AthenaWorkGroup",
				"",
			},
			wantErr: true,
		},
		{
			name: "AthenaWorkGroup is valid",
			args: args{
				"AthenaWorkGroup",
				"This@is-valid-",
			},
			wantErr: false,
		},
		{
			name: "AthenaWorkGroup is not valid",
			args: args{
				"AthenaWorkGroup",
				"This@is-not#valid-",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAthenaWorkGroup(tt.args.attribute, tt.args.athenaWorkGroup); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAthenaWorkGroup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLogResourcePrefix(t *testing.T) {
	type args struct {
		attribute         string
		logResourcePrefix string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty LogResourcePrefix is invalid",
			args: args{
				"LogResourcePrefix",
				"",
			},
			wantErr: true,
		},
		{
			name: "LogResourcePrefix is valid",
			args: args{
				"LogResourcePrefix",
				"abcdefgh",
			},
			wantErr: false,
		},
		{
			name: "LogResourcePrefix is not valid",
			args: args{
				"LogResourcePrefix",
				"aBcDeFgH",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLogResourcePrefix(tt.args.attribute, tt.args.logResourcePrefix); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLogResourcePrefix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateRoleARN(t *testing.T) {
	type args struct {
		attribute string
		roleARN   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty RoleARN is invalid",
			args: args{
				"RoleARN",
				"",
			},
			wantErr: true,
		},
		{
			name: "AWSAccountID in RoleARN is to long",
			args: args{
				"RoleARN",
				"arn:aws:iam::0197577444913:role/CustomerManagedEndpoint",
			},
			wantErr: true,
		},
		{
			name: "RoleARN is valid",
			args: args{
				"RoleARN",
				"arn:aws:iam::197577444913:role/CustomerManagedEndpoint",
			},
			wantErr: false,
		},
		{
			name: "RoleARN is not valid",
			args: args{
				"RoleARN",
				"xyzarn:aws:iam::197577444913:role/CustomerManagedEndpoint",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateRoleARN(tt.args.attribute, tt.args.roleARN); (err != nil) != tt.wantErr {
				t.Errorf("ValidateRoleARN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAWSAccountID(t *testing.T) {
	type args struct {
		attribute    string
		AWSAccountID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty AWSAccountID is invalid",
			args: args{
				"AWSAccountID",
				"",
			},
			wantErr: true,
		},
		{
			name: "AWSAccountID is to short",
			args: args{
				"AWSAccountID",
				"019757744",
			},
			wantErr: true,
		},
		{
			name: "AWSAccountID is to long",
			args: args{
				"AWSAccountID",
				"0197577444913",
			},
			wantErr: true,
		},
		{
			name: "AWSAccountID is valid",
			args: args{
				"AWSAccountID",
				"197577444913",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAWSAccountID(tt.args.attribute, tt.args.AWSAccountID); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAWSAccountID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateLogDestination(t *testing.T) {
	type args struct {
		attribute      string
		logDestination string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty LogDestination is invalid",
			args: args{
				"LogDestination",
				"",
			},
			wantErr: true,
		},
		{
			name: "LogDestination is valid",
			args: args{
				"LogDestination",
				"this-1-is-valid-",
			},
			wantErr: false,
		},
		{
			name: "LogDestination has uppercase",
			args: args{
				"LogDestination",
				"this-1-is-NOT-valid",
			},
			wantErr: true,
		},
		{
			name: "LogDestination has underscore",
			args: args{
				"LogDestination",
				"this-1-is_also_invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateLogDestination(tt.args.attribute, tt.args.logDestination); (err != nil) != tt.wantErr {
				t.Errorf("ValidateLogDestination() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNameStrict(t *testing.T) {
	type args struct {
		attribute string
		name      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty Name is invalid",
			args: args{
				"Name",
				"",
			},
			wantErr: true,
		},
		{
			name: "Name is valid",
			args: args{
				"Name",
				"this-1-is_VERY_valid",
			},
			wantErr: false,
		},
		{
			name: "Name has odd character",
			args: args{
				"Name",
				"this-1-is-NOT-v@lid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNameStrict(tt.args.attribute, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNameStrict() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
