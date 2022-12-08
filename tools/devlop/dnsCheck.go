package devlop

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

var ipList = []string{
	"114.114.114.114", "114.114.115.115", "223.5.5.5", "223.6.6.6", "180.76.76.76", "119.29.29.29", "182.254.116.116", "1.2.4.8", "210.2.4.8", "117.50.11.11", "52.80.66.66", "101.226.4.6", "218.30.118.6", "123.125.81.6", "140.207.198.6", "1.1.1.1", "1.0.0.1", "8.8.8.8", "8.8.4.4", "9.9.9.9", "185.222.222.222", "185.184.222.222", "208.67.222.222", "208.67.220.220", "199.91.73.222", "178.79.131.110", "61.132.163.68", "202.102.213.68", "219.141.136.10", "219.141.140.10", "61.128.192.68", "61.128.128.68", "218.85.152.99", "218.85.157.99", "202.100.64.68", "61.178.0.93", "202.96.128.86", "202.96.128.166", "202.96.134.133", "202.96.128.68", "202.103.225.68", "202.103.224.68", "202.98.192.67", "202.98.198.167", "222.88.88.88", "222.85.85.85", "219.147.198.230", "219.147.198.242", "202.103.24.68", "202.103.0.68", "222.246.129.80", "59.51.78.211", "218.2.2.2", "218.4.4.4", "61.147.37.1", "218.2.135.1", "202.101.224.69", "202.101.226.68", "219.148.162.31", "222.74.39.50", "219.146.1.66", "219.147.1.66", "218.30.19.40", "61.134.1.4", "202.96.209.133", "116.228.111.118", "202.96.209.5", "180.168.255.118", "61.139.2.69", "218.6.200.139", "219.150.32.132", "219.146.0.132", "222.172.200.68", "61.166.150.123", "202.101.172.35", "61.153.177.196", "61.153.81.75", "60.191.244.5", "123.123.123.123", "123.123.123.124", "202.106.0.20", "202.106.195.68", "221.5.203.98", "221.7.92.98", "210.21.196.6", "221.5.88.88", "202.99.160.68", "202.99.166.4", "202.102.224.68", "202.102.227.68", "202.97.224.69", "202.97.224.68", "202.98.0.68", "202.98.5.68", "221.6.4.66", "221.6.4.67", "202.99.224.68", "202.99.224.8", "202.102.128.68", "202.102.152.3", "202.102.134.68", "202.102.154.3", "202.99.192.66", "202.99.192.68", "221.11.1.67", "221.11.1.68", "210.22.70.3", "210.22.84.3", "119.6.6.6", "124.161.87.155", "202.99.104.68", "202.99.96.68", "221.12.1.227", "221.12.33.227", "202.96.69.38", "202.96.64.68", "221.131.143.69", "112.4.0.55", "211.138.180.2", "211.138.180.3", "218.201.96.130", "211.137.191.26",
}

var addrList = []string{
	"web.medlinker.com", "app.v3.medlinker.com", "web.v3.medlinker.com", "www.medlinker.com", "android.medlinker.com", "passport.medlinker.com", "ylt.medlinker.com", "h5.medlinker.com", "app.medlinker.com", "openapi.medlinker.com", "open.medlinker.com", "im-gateway.medlinker.com", "im-api.medlinker.com", "quiz.medlinker.com", "search.svr.medlinker.com", "im-wss.medlinker.com", "quiz-mini-svr.medlinker.com", "ph.medlinker.com", "cloudpharmacy.medlinker.com", "tim.medlinker.com", "wzapi.medlinker.com", "confluence.medlinker.com", "yax.medlinker.com", "inquiry.medlinker.com", "rank.medlinker.com", "hc-svr.medlinker.com", "hc-user-svr.medlinker.com", "hc-panel-svr.medlinker.com", "hc-medication-svr.medlinker.com", "yjgwyy-doctor.medlinker.com", "yjgwyy-patient.medlinker.com", "yjgwyy-panel.medlinker.com", "hc-patient.medlinker.com", "inquiry-panel-svr.medlinker.com", "inquiry-panel.medlinker.com", "rn.medlinker.com", "med-data.medlinker.com", "diancan.medlinker.com", "inquiry-openapi.medlinker.com", "reporthelper.medlinker.com", "inquiry-h5.medlinker.com", "med-payment.medlinker.com", "med-files.medlinker.com", "yjgwyy-doctor-h5.medlinker.com", "med-order.medlinker.com", "cloud-pharmacy.medlinker.com", "doctor-app.medlinker.com", "emall.medlinker.com", "patient-medication.medlinker.com", "med-wechat.medlinker.com", "med-cdm.medlinker.com", "business-new.medlinker.com", "hc-svr-sc.medlinker.com", "hc-panel-svr-sc.medlinker.com", "hc-user-svr-sc.medlinker.com", "cdhesyy-patient.medlinker.com", "cdhesyy-manage.medlinker.com", "cdhesyy-doctor.medlinker.com", "cdhesyy-doctor-h5.medlinker.com", "hc-patient-sc.medlinker.com", "emall-api.medlinker.com", "monitor.medlinker.com", "hc-user-hz.medlinker.com", "hc-svr-hz.medlinker.com", "hc-panel-hz.medlinker.com", "hzsdermyy-doctor-h5.medlinker.com", "hzsdermyy-doctor.medlinker.com", "hzsdermyy-manage.medlinker.com", "hzsdermyy-patient.medlinker.com", "ih-openapi.medlinker.com", "upgrade.medlinker.com", "pmm-server.medlinker.com", "emall-cron.medlinker.com", "wz-panel.medlinker.com", "task.medlinker.com", "hc-user-xj.medlinker.com", "hc-svr-xj.medlinker.com", "hc-panel-xj.medlinker.com", "wlmqdsyy-doctor-h5.medlinker.com", "wlmqdsyy-doctor.medlinker.com", "wlmqdsyy-panel.medlinker.com", "wlmqdsyy-patient.medlinker.com", "business.medlinker.com", "m.medlinker.com", "hc-patient-xj.medlinker.com", "team.medlinker.com", "med-questionnaire.medlinker.com", "med-panel.medlinker.com", "med-cooperate-medication.medlinker.com", "dt.medlinker.com", "med-broker-app.medlinker.com", "med-cloud-wallet.medlinker.com", "workstation-api.medlinker.com", "med-doctor-grpc-pprof.medlinker.com", "prod.external.medlinker.com", "med-delay-queue.medlinker.com", "med-web-thirdparty.medlinker.com", "med-clinic-business.medlinker.com", "external-prod.medlinker.com", "general-interface.medlinker.com", "im-interface.medlinker.com", "web-monitor-panel.medlinker.com", "web-monitor.medlinker.com", "department-meeting.medlinker.com", "cdm-patient-tools.medlinker.com", "xim-pprof.medlinker.com", "broker-interface.medlinker.com", "interview-web.medlinker.com", "med-interview.medlinker.com", "med-ih-listener.medlinker.com", "amp.medlinker.com", "pmp.medlinker.com", "med-gateway-admin.medlinker.com", "med-order-logistics.medlinker.com", "wechat-mp.medlinker.com", "ab-admin.medlinker.com", "ab-web.medlinker.com", "cdhesyy-agent-h5.medlinker.com", "cdhesyy-agent-manage.medlinker.com", "med-captcha.medlinker.com", "med-toolchain.medlinker.com", "med-one-out.medlinker.com", "msso.medlinker.com", "msso-oa.medlinker.com", "med-oauth.medlinker.com", "med-sso-panel.medlinker.com", "med-sso-auth.medlinker.com", "supplier-chain-panel.medlinker.com", "med-sso.medlinker.com", "message-service.medlinker.com", "open-platform-panel.medlinker.com", "open-platform.medlinker.com", "med-ums-admin.medlinker.com", "chronic-panel.medlinker.com", "inspect-panel.medlinker.com", "patitent-cross-platform.medlinker.com", "service-panel.medlinker.com", "payment-center.medlinker.com", "doctor-patient-operation.medlinker.com", "ordering-system-admin.medlinker.com", "ordering-system-h5.medlinker.com", "web-mall.medlinker.com", "revenue-panel.medlinker.com", "meal-admin.medlinker.com", "meal-web.medlinker.com", "med-ums-web.medlinker.com", "doctor-panel.medlinker.com", "patient-panel.medlinker.com", "public-chart.medlinker.com", "huafang-panel.medlinker.com", "app-update-panel.medlinker.com", "med-rn-web.medlinker.com", "med-meal-admin.medlinker.com", "med-meal-web.medlinker.com", "supply-chain-statement.medlinker.com", "r.medlinker.com", "monitor-config-panel.medlinker.com", "scale-panel.medlinker.com", "med-abtest-admin.medlinker.com", "qw-panel.medlinker.com", "pop.medlinker.com", "hotupdate.medlinker.com", "rn-hotupdate.medlinker.com", "login.medlinker.com", "online-sale-panel.medlinker.com", "broker-panel.medlinker.com", "chart-data.medlinker.com", "hao123.medlinker.com", "med-gov-push.medlinker.com", "medrax.medlinker.com", "pmo.medlinker.com", "app-dialog-content.medlinker.com", "med-rn-almond.medlinker.com", "medhybrid.medlinker.com", "med-calendar.medlinker.com", "one-data-panel.medlinker.com", "bi-board-ding.medlinker.com", "osteology-panel.medlinker.com", "med-pop-outside.medlinker.com", "golang-demo.medlinker.com.medlinker.com", "one-data.medlinker.com", "med-one-example.medlinker.com", "chronic-panel-v2.medlinker.com", "tumor-adverse-reaction-panel.medlinker.com", "med-one-example1.medlinker.com", "med-one-example2.medlinker.com", "med-one-example3.medlinker.com", "twy.medlinker.com", "m-twy.medlinker.com", "m-zhkf.medlinker.com", "zhkf.medlinker.com", "calculator.doctorword.com.medlinker.com", "med-wallet-withdraw.medlinker.com", "med-wallet-go.medlinker.com", "dietitian-panel.medlinker.com", "ten-migrate-test1.medlinker.com", "ten-migrate-test2.medlinker.com", "disease-sop-pprof.medlinker.com", "course-panel.medlinker.com", "med-drugstore-for-apply.medlinker.com", "broker-m.medlinker.com", "notice.medlinker.com", "risk-panel.medlinker.com", "pay.medlinker.com",
}

var wantIpList = []string{
	"47.111.100.76",
	"101.37.45.130",
	"120.55.137.110",
}

var wg = new(sync.WaitGroup)

func Dns() {
	t := time.Now()
	for _, ip := range ipList {
		for _, addr := range addrList {
			wg.Add(1)
			go func(ip, addr string) {
				if err := do(ip, addr); err != nil {
					println(err.Error())
				}
			}(ip, addr)
		}
	}
	wg.Wait()
	duration := time.Since(t)
	fmt.Printf("一共用了：%d 秒", duration/time.Second)
}

func do(addr, host string) (err error) {
	defer wg.Done()
	resolver := net.Resolver{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return net.Dial("tcp", addr)
		},
		StrictErrors: true,
	}
	ns, err := resolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return err
	}
	for _, n := range ns {
		nstr := n.String()
		for _, s := range wantIpList {
			if nstr == s {
				return
			}
		}
		return errors.New(fmt.Sprintf("服务商ip (%s), 医联地址 (%s), 解出的ns (%s)", addr, host, nstr))
	}
	return
}
