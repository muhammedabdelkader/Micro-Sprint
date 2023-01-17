package main

import (
	"fmt"
)

type Component struct {
	    ID uuid.UUID
		    Name string
			    systemID string
				Spoofing struct {
					        Risks []string
							        Remediation []string
									    
				}
				Tampering struct {
					        Risks []string
							        Remediation []string
									    
				}
				Repudiation struct {
					        Risks []string
							        Remediation []string
									    
				}
				InformationDisclosure struct {
					        Risks []string
							        Remediation []string
									    
				}
				DenialOfService struct {
					        Risks []string
							        Remediation []string
									    
				}
				ElevationOfPrivilege struct {
					        Risks []string
							        Remediation []string
									    
				}
				
}


type Component struct {
	ID       uuid.UUID
	Name     string
	Threats  []string
	systemID string
	controls map[string][]string
}
/*
func main() {
	// Define the system's components
	component := Component{
		uuid.NewV4(),
		"Authentication system",
		[]string{"Spoofing", "Tampering", "Repudiation"},
		"System1",
		map[string][]string{
			"Spoofing":    []string{"Implement multi-factor authentication", "Monitor login attempts for unusual activity"},
			"Tampering":   []string{"Use tamper-evident logs", "Implement data validation on input"},
			"Repudiation": []string{"Implement session management", "Implement digital signature on transactions"},
		},
	}

	// Print the component's associated threats and controls
	fmt.Printf("ID: %s\n", component.ID)
	fmt.Printf("%s (systemID: %s):\n", component.Name, component.systemID)
	for _, threat := range component.Threats {
		fmt.Printf("\t- %s\n", threat)
		fmt.Printf("\t\tRemediation Suggestions: %s\n", component.controls[threat])

	}

}*/

func main() {
	    // Define the system's components
		component := Component{
			        ID: uuid.NewV4(), 
					        Name: "Authentication system", 
							        systemID: "System1", 
									Spoofing: &Threat{
										            Risks: []string{"Impersonation", "Session hijacking"},
													            Remediation: []string{"Implement multi-factor authentication", "Monitor login attempts for unusual activity"}
																        
									},
									Tampering: &Threat{
										            Risks: []string{"Data tampering", "Configuration manipulation"},
													            Remediation: []string{"Use tamper-evident logs", "Implement data validation on input"}
																        
									},
									Repudiation: &Threat{
										            Risks: []string{"Denial of service", "Unauthorized access"},
													            Remediation: []string{"Implement session management", "Implement digital signature on transactions"}
																        
									},
									InformationDisclosure: &Threat{
										            Risks: []string{},
													            Remediation: []string{}
																        
									},
									DenialOfService: &Threat{
										            Risks: []string{},
													            Remediation: []string{}
																        
									},
									ElevationOfPrivilege: &Threat{
										            Risks: []string{},
													            Remediation: []string{}
																        
									},
									    
		}
		    fmt.Println("Component Name: ", component.Name)
			    fmt.Println("Spoofing Risks: ", component.Spoofing.Risks)
				    fmt.Println("Spoof
					")
					fmt.Println("Tampering Risks: ", component.Tampering.Risks)
					    fmt.Println("Tampering Remediation: ", component.Tampering.Remediation)
						    fmt.Println("Repudiation Risks: ", component.Repudiation.Risks)
							    fmt.Println("Repudiation Remediation: ", component.Repudiation.Remediation)
								    fmt.Println("Information Disclosure Risks: ", component.InformationDisclosure.Risks)
									    fmt.Println("Information Disclosure Remediation: ", component.InformationDisclosure.Remediation)
										    fmt.Println("Denial of Service Risks: ", component.DenialOfService.Risks)
											    fmt.Println("Denial of Service Remediation: ", component.DenialOfService.Remediation)
												    fmt.Println("Elevation of Privilege Risks: ", component.ElevationOfPrivilege.Risks)
													    fmt.Println("Elevation of Privilege Remediation: ", component.ElevationOfPrivilege.Remediation)

}
