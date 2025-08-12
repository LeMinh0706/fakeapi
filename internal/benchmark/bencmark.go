package benchmark

import "github.com/gin-gonic/gin"

var benchmarkData []gin.H = []gin.H{
	{
		"candidate_id":  "e3e54b26-f1ce-44b0-8d86-cb27e5237bf3",
		"email":         "leminhken124356@gmail.com",
		"applied":       "backend",
		"overall_score": 34.65,
		"score": gin.H{
			"certificate": gin.H{
				"total_score": 0.0,
			},
			"language": gin.H{
				"total_score": 35.0,
				"detailed_score": gin.H{
					"english":         35.0,
					"other_languages": 0.0,
				},
			},
			"education": gin.H{
				"total_score": 77.0,
				"detailed_score": gin.H{
					"degree":     19.0,
					"university": 18.0,
					"major":      20.0,
					"gpa":        20.0,
				},
			},
			"experience": gin.H{
				"total_score": 0.0,
				"detailed_score": gin.H{
					"job_relevance":       0.0,
					"stability":           0.0,
					"predicted_level":     "Fresher",
					"matching_job_titles": []string{},
				},
			},
			"skills": gin.H{
				"total_score": 59.5,
				"detail_score": gin.H{
					"python - ai":                     67.7,
					"java - csm":                      58.6,
					"dart - safe":                     57.1,
					"c# - csm":                        61.8,
					"spring boot - agile":             55.8,
					"mvc - agile":                     58.6,
					"firebase - agile":                51.6,
					"sql server - enterprise systems": 67.9,
					"postman - pmp":                   61.8,
					"expo go - agile":                 52.9,
					"visual studio code - software development": 71.8,
					"flutter - agile": 55.2,
					"nodejs - csm":    56.1,
					"html - csm":      56.2,
					"css - csm":       66.9,
					"javascript - ai": 56.4,
					"bootstrap - csm": 54.8,
				},
			},
		},
	},
	{
		"candidate_id":  "7575acae-c9df-4919-a5a4-648abb481f46",
		"email":         "nguyenhuuluanit69@gmail.com",
		"applied":       "fullstack",
		"overall_score": 49.763125,
		"score": gin.H{
			"certificate": gin.H{
				"total_score": 100.0,
			},
			"language": gin.H{
				"total_score": 75.0,
				"detailed_score": gin.H{
					"english":         65.0,
					"other_languages": 10.0,
				},
			},
			"education": gin.H{
				"total_score": 76.6875,
				"detailed_score": gin.H{
					"degree":     19.0,
					"university": 18.0,
					"major":      20.0,
					"gpa":        19.6875,
				},
			},
			"experience": gin.H{
				"total_score": 0.0,
				"detailed_score": gin.H{
					"job_relevance":       0.0,
					"stability":           0.0,
					"predicted_level":     "Fresher",
					"matching_job_titles": []string{},
				},
			},
			"skills": gin.H{
				"total_score": 56.7,
				"detail_score": gin.H{
					"flutter - agile":                     55.2,
					"express.js - ai":                     50.0,
					"spring boot - agile":                 55.8,
					"asp.net core - software development": 50.6,
					"c# - csm":                            61.8,
					"java - csm":                          58.6,
					"dart - safe":                         57.1,
					"c/c++ - csm":                         58.1,
					"javascript - ai":                     56.4,
					"typescript - csm":                    51.7,
					"sql server - enterprise systems":     67.9,
					"docker - agile":                      50.2,
					"postman - pmp":                       61.8,
					"figma - ai":                          54.6,
					"jira - ai":                           53.4,
					"notion - csm":                        57.2,
					"git - software development":          56.6,
					"solid - safe":                        62.0,
					"oop - ai":                            63.5,
					"design pattern - agile":              58.9,
					"firebase - agile":                    51.6,
					"gcp - pmp":                           61.8,
					"jwt - csm":                           51.9,
					"agora - ai":                          54.3,
				},
			},
		},
	},
	{
		"candidate_id":  "15cc89ef-e0a7-4d88-8792-69ebbc02ba76",
		"email":         "phamnamduong583@gmail.com",
		"applied":       "frontend",
		"overall_score": 44.325625,
		"score": gin.H{
			"certificate": gin.H{
				"total_score": 100.0,
			},
			"language": gin.H{
				"total_score": 35.0,
				"detailed_score": gin.H{
					"english":         35.0,
					"other_languages": 0.0,
				},
			},
			"education": gin.H{
				"total_score": 77.4375,
				"detailed_score": gin.H{
					"degree":     19.0,
					"university": 18.0,
					"major":      20.0,
					"gpa":        20.4375,
				},
			},
			"experience": gin.H{
				"total_score": 0.0,
				"detailed_score": gin.H{
					"job_relevance":       0.0,
					"stability":           0.0,
					"predicted_level":     "Fresher",
					"matching_job_titles": []string{},
				},
			},
			"skills": gin.H{
				"total_score": 58.2,
				"detail_score": gin.H{
					"tensorflow - ai":                          56.0,
					"pytorch - ai":                             59.6,
					"opencv - safe":                            52.7,
					"keras - ai":                               52.3,
					"python - ai":                              67.7,
					"c++ - csm":                                62.0,
					"java - csm":                               58.6,
					"sql - data-driven transformation":         57.8,
					"oop - ai":                                 63.5,
					"cnns - machine learning":                  60.6,
					"gans - ai":                                64.4,
					"linear regression - machine learning":     52.8,
					"decision trees - machine learning":        67.3,
					"clustering algorithms - machine learning": 69.9,
					"deep learning architectures - machine learning": 65.2,
					"computer vision - ai":                           63.2,
					"android - software development":                 56.4,
					"camerax - ai":                                   51.3,
					"retrofit - agile":                               55.6,
					"bitmap processing - data-driven transformation": 50.0,
					"lstm - machine learning":                        65.5,
					"fastapi - ai":                                   52.4,
					"numpy - ai":                                     51.1,
					"uvicorn - safe":                                 50.6,
					"json - machine learning":                        56.0,
					"videoutils - safe":                              54.0,
					"javascript - ai":                                56.4,
					"node.js - csm":                                  51.6,
					"bootstrap - csm":                                54.8,
					"css - csm":                                      66.9,
					"scss - csm":                                     60.2,
					"c# - csm":                                       61.8,
					"html - csm":                                     56.2,
					"jquery - machine learning":                      56.3,
				},
			},
		},
	},
	{
		"candidate_id":  "37eadf14-d329-4509-b5cd-8baade5e3efd",
		"email":         "camtu28082003@gmail.com",
		"applied":       "frontend",
		"overall_score": 61.09,
		"score": gin.H{
			"certificate": gin.H{
				"total_score": 100.0,
			},
			"language": gin.H{
				"total_score": 35.0,
				"detailed_score": gin.H{
					"english":         35.0,
					"other_languages": 0.0,
				},
			},
			"education": gin.H{
				"total_score": 74.0,
				"detailed_score": gin.H{
					"degree":     19.0,
					"university": 25.0,
					"major":      20.0,
					"gpa":        10.0,
				},
			},
			"experience": gin.H{
				"total_score": 57.0,
				"detailed_score": gin.H{
					"job_relevance":   40.0,
					"stability":       17.0,
					"predicted_level": "Middle",
					"matching_job_titles": []string{
						"Penetration Tester",
						"Bug Bounty Hunter",
					},
				},
			},
			"skills": gin.H{
				"total_score": 58.8,
				"detail_score": gin.H{
					"penetration testing - safe":                     55.1,
					"information gathering - enterprise systems":     58.1,
					"vulnerability detection - software development": 61.6,
					"exploitation - agile":                           66.9,
					"burp suite - pmp":                               51.9,
					"acunetix - pmi-acp":                             64.4,
					"sql map - data-driven transformation":           59.1,
					"ffuf - safe":                                    56.1,
					"cve exploitation - agile":                       60.6,
					"cyber attacks - digital transformation":         55.9,
					"dos attacks - csm":                              57.0,
					"self-developed tools - software development":    67.0,
					"social engineering - enterprise systems":        61.1,
					"information security - enterprise systems":      58.7,
					"vulnerability finding - safe":                   56.1,
					"reflectedxss - csm":                             54.4,
					"wordpress - csm":                                61.1,
					"botnet - ai":                                    53.4,
				},
			},
		},
	},
}

// func BenchmarkResult() ([]gin.H, []gin.H) {
// 	var res1, res2 []gin.H

// 	results, ok := benchmarkData["benchmark_results"].([]gin.H)
// 	if !ok {
// 		return res1, res2
// 	}

// 	for i, v := range results {
// 		if i%2 == 0 {
// 			res1 = append(res1, v)
// 		} else {
// 			res2 = append(res2, v)
// 		}
// 	}
// 	return res1, res2
// }
