package common

// func (ks *kubeService) GetVSphereCluster(ctx context.Context, options ...GetVSphereClusterOptions) (*GetVSphereClusterReturn, error) {
// 	option := TakeFirstOrDefaults(GetVSphereClusterOptions{}, options)
// 	_ = option

// 	clusterName := option.ClusterName

// 	if clusterName == "" {
// 		if option.ClusterID == "" {
// 			return nil, errors.New("cluster name or cluster id is required")
// 		}

// 		// Get cluster name from config map
// 		configMap, err := ks.GetConfigMap(ctx, GetConfigMapOption{
// 			ConfigMapName: GetConfigMapName(option.ClusterID),
// 		})
// 		if err != nil {
// 			return nil, err
// 		}

// 		clusterName = configMap.ClusterName
// 	}

// 	clusterAPICtl, err := ClusterCtlWrapper.New(ClusterCtlWrapper.NewOptions{
// 		ClusterName: clusterName,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	status, err := clusterAPICtl.GetClusterStatus()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var clusterLogs []*LogEntry
// 	if status.Status.Phase != ClusterStatuses.Provisioned.String() {
// 		if _clusterLogs, err := ks.GetVSphereClusterLogs(ctx, GetVsphereClusterLogsOptions{
// 			ClusterID:   option.ClusterID,
// 			ClusterName: option.ClusterName,
// 			MaxLines:    option.MaxLogLines,
// 		}); err != nil {
// 			fmt.Println(fmt.Errorf("failed to get cluster logs: %v", err))
// 		} else {
// 			clusterLogs = _clusterLogs.LogEntries
// 		}
// 	}

// 	return &GetVSphereClusterReturn{
// 		ClusterName: clusterName,
// 		Status:      status.Status,
// 		Logs:        clusterLogs,
// 	}, nil
// }

// func (ks *kubeService) GetVSphereClusterLogs(ctx context.Context, options ...GetVsphereClusterLogsOptions) (*GetVsphereClusterLogsReturn, error) {
// 	option := utils.TakeFirstOrDefaults(GetVsphereClusterLogsOptions{}, options)

// 	if option.MaxLines == 0 {
// 		option.MaxLines = 100
// 	}

// 	clientset := ks.clientset

// 	// Define the namespace and the label selector for the pod name
// 	namespace := "capv-system"
// 	podNamePrefix := "capv-controller-manager"

// 	// Get the CAPZ pod in capz-system namespace
// 	// List the pods in the namespace with a field selector for the name prefix
// 	podList, err := clientset.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{
// 		LabelSelector: fmt.Sprintf("control-plane=%s", podNamePrefix), // Assuming pods are labeled by app
// 	})
// 	if err != nil {
// 		fmt.Printf("Failed to list pods: %w\n", err)
// 		return nil, fmt.Errorf("failed to list pods: %w", err)
// 	}

// 	// Check if any pods matched the query
// 	if len(podList.Items) == 0 {
// 		fmt.Println("No pods found with the given name prefix.")
// 		return nil, fmt.Errorf("no pods found")
// 	}

// 	// Select the first pod
// 	selectedPod := podList.Items[0].Name
// 	fmt.Printf("Selected pod: %s\n", selectedPod)

// 	scanLines := int64(1000)

// 	// Get the logs from the pod
// 	// Request logs for the selected pod
// 	req := clientset.CoreV1().Pods(namespace).GetLogs(selectedPod, &v2.PodLogOptions{
// 		TailLines: &scanLines,
// 	})
// 	podLogs, err := req.Stream(ctx)
// 	if err != nil {
// 		fmt.Printf("Failed to get logs: %v\n", err)
// 		return nil, fmt.Errorf("failed to get logs: %w", err)
// 	}
// 	defer func(podLogs io.ReadCloser) {
// 		err := podLogs.Close()
// 		if err != nil {
// 			fmt.Printf("Failed to close logs: %v\n", err)
// 		}
// 	}(podLogs)

// 	clusterNameSelector := fmt.Sprintf("%s-control-plane", option.ClusterName)

// 	// Filter log entries based on cluster name
// 	logEntries, _ := utils.ParseLogs(podLogs, utils.ParseLogsOptions{
// 		MaxLines:  scanLines,
// 		ErrorOnly: true,
// 		FilterBy: []utils.KeyValue{
// 			{
// 				Key:   "name",
// 				Value: clusterNameSelector,
// 			},
// 		},
// 	})

// 	// Take only last 100 log entries
// 	if len(logEntries) > int(option.MaxLines) {
// 		logEntries = logEntries[len(logEntries)-int(option.MaxLines):]
// 	}

// 	return &GetVsphereClusterLogsReturn{LogEntries: logEntries}, nil
// }

// type GetVsphereClusterLogsReturn struct {
// 	Logs       []string
// 	LogEntries []*utils.LogEntry
// }

// type GetVsphereClusterLogsOptions struct {
// 	ClusterID   string
// 	ClusterName string
// 	MaxLines    int64
// }
