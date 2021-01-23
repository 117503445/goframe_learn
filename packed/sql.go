package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GDoaDQPYkACYgycDMlFqYklqXrFhTl6+gh2aAgrA+NDL56EN+Gx/o9iBPapP7Q/08o5oSFz9fo9XjtWvGVQnxJcXHVDNP8Ws25e7L6grtX/fvlqxyetDDWZnbMt3GBu5bPfiRsmMIdaZxxUkvln/27+tU87a/Qmmfm0LfF8a/YtZ21r1O0mK4+Pr71YtTT91xxwrX9y0IZHrLlAsMJgDpuAj/s9m+ad585oP3i1tNFBquH6kUdiFhu3Bp9+btKmdXzpoSCegDPPE94//rGOufb/tfktF5aKZQZdaZwlk3HdwOFt0WW1kOQbjqdOr3RoPiC88PktpZIDogvFfyWvrwiqYru0cWFZb/qVkhlHz/zlkbguXTmfwVgw+XzjxZIi6Sp+xQzdqSul3NkuKuiwM5tcNli0S3//7Us7VW4e2s7+77bcOomXK9aw575xNMxy49wlttB3xxMew6tTAmTXL2azu1G4+u3aOd23WSsYZ56fypA5OaDFasLJlFXx77ZIvlJ4nxJVPzXbVjM6rW67tNrrFRNb/1kIJtdPqNSfYT6D7Z3MssNnXgusFshSfOygYX7d+L+v6a4Mx0A9Z8eJIjuVToo+ZVF56RSot+Rz8tM3ayb80nonLdJi/U29Neitr/upxj6xYv2OvyknqvjbXQQ+KqQ0eDybaHooW+/mnqpmf4V7Dz2seyOCnlu1ddTIsn0s2hymONX8wk1zXrPbZn8y12d/jJnwiJ3RdHLUnJUtn1SuXF1/xuhPL3cbb4qHaYzw/y5zHy4rqROCElFzBHvdF72bui6v004u7Eum+ZrQ0Pg2Uc/7kmFC2n8DuarOm70tOnxtU47gxDkuyidXnXroM6HV7vtmh0OmG3Ne7c9+xDBjRVt8XaCFseH8/7PY3l2vEvyuJNLI/F1J4o+AY+A0lZVfDN1vav+wvB5kc6r+/NYppXcn71C/vLZk1/8Wy7t7p54KsAg6MnGNTKx03e+Nt1aGse/9f64vtfH5E8fvt9RCo/+q/3SblTh1xmN1szl3P4bdfn/ny7PysyfqKtnf/l2tev3TzvuMDAz//wd4s3M0Rxv9usLEwDCPl4EBlkEYGHrQMggXSgYBZ4qXXjwJIAOQlQV4MzKJMCPyGLLhoDwGA9saQSSBHIcwDLt7IECA4b+jIjMDdtexsoGUMDEwMfQxMDAkMYN4gAAAAP//wzH//wYEAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
