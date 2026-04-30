00:00:00,380 --> 00:00:02,940 
Ja, sonst, sonst hab ich halt von dir-

00:00:02,949 --> 00:00:04,210 
So, jetzt wird's, glaub ich, recorded.

00:00:04,280 --> 00:00:12,760 
Genau. Ja, sonst hab ich halt von dir nur die Likert-Evaluierung. Ist auch fein. Genau.
Kommen wir mal zu den organisatorischen Fragen. [lacht]

00:00:12,800 --> 00:00:13,080 
Mhm.

00:00:15,140 --> 00:00:16,640 
Was ist deine aktuelle Position?

00:00:18,000 --> 00:00:20,720 
Äh, Senior Lecturer an der Hochschule.

00:00:22,400 --> 00:00:27,440 
Ähm, wie lang machst du deine aktuelle Rolle Schrägstrich wie lang bist du in der IT?

00:00:28,720 --> 00:00:36,720 
Wie lang ich in der IT bin? Äh, gut neun Jahre. Außer wenn man jetzt,
wenn man jetzt mein Studium ausschließt.

00:00:37,050 --> 00:00:44,700 
[lacht] Verstehe. Was ist,
hast du eine tägliche Exposure zu Application Configuration und Secret Management?

00:00:46,180 --> 00:01:06,820 
Application Configuration, äh, hatte ich sehr viel Exposure vor allem, äh, Secret Management, äh,
und von Secret Management hatte ich auch sehr viel bei meiner, äh, letzten Arbeit, wo ich, äh,

00:01:07,540 --> 00:01:16,370 
Wunderbar. Warst du mit deinen Aproaches, die du vorgefunden hast, satisfied?
Erst recht im Kontext auf hybride Szenarien?

00:01:18,880 --> 00:01:30,400 
Äh, jein.
Also oftmals waren die Approaches von Anfang an schon lacking und dementsprechend mussten sie

00:01:33,120 --> 00:01:35,700 
Verstehe. Hast du schon mal den External Secret Operator verwendet?

00:01:37,380 --> 00:01:38,340 
Äh, ja.

00:01:40,360 --> 00:02:15,180 
Ist wunderbar. Das heißt, da brauchen wir nicht unendlich viel erklären. Gut,
dann kommen wir auch schon in den spannenderen Teil, was ich jetzt genau gemacht hab.

00:02:15,190 --> 00:02:17,700 
Äh, die ist schon oben.

00:02:17,760 --> 00:02:29,120 
Ah, da der große Button, der ist mir zu groß gewesen. Genau.
Was ein centralized Secret Store ist und was,

00:02:29,640 --> 00:02:29,880 
Mhm.

00:02:30,660 --> 00:02:49,720 
Wunderbar. Genau, dann der External Secret Operator noch mal eine kurze Zusammenfassung:
Ist ein Tool, der redet für dich mit dem Vault und kreiert dir daraus Kubernetes-native Secrets,

00:02:50,109 --> 00:02:50,109 
Mhm.

00:02:50,120 --> 00:03:22,060 
Und das ist ja eigentlich ein supercooles Konzept per se konzeptionell,
aber du hast halt ein bissl das Problem, dass du so was auf Linux nicht hast.

00:03:22,540 --> 00:03:22,620 
Mhm.

00:03:22,640 --> 00:04:09,920 
Genau. Und Dapr ist sozusagen die Distributed Application Runtime, kann gut sein,
von der gehört hast. Das ist so ein Building-Block-System,

00:04:09,960 --> 00:04:21,820 
Das heißt, nur dass ich dich verstanden hab, das heißt, der managed dann mehr oder weniger das, äh,
das, also er handelt das File-Management,

00:04:22,380 --> 00:04:33,840 
Genau, also dem sozusagen gibst du so wie, als wenn du halt eine Applikation laufen hast,
die von ETC App, die die B-Environment-Parameter sitzt, äh, frisst,

00:04:33,930 --> 00:04:34,080 
Mhm.

00:04:34,270 --> 00:04:44,640 
Und der macht dir dann dieses File und wenn du die Time to leave kannst du halt einstellen und der
rolliert dir das halt durchgängig, dass du sozusagen immer dynamic-secret-mäßig das hast,

00:04:44,680 --> 00:04:44,750 
Mhm.

00:04:45,240 --> 00:05:05,900 
Genau, und konzeptionell heißt das, du hast den Git Repository,
wo du halt mit CI/CD entweder über kubectl apply oder mit einem Curl auf die richtigen Pfade die

00:05:05,940 --> 00:05:07,000 
Mhm.

00:05:07,080 --> 00:06:43,900 
Genau. Jetzt genauer zu meinem Solution Approach: Was habe ich gemacht? Ich hab dann,
ich hab sozusagen die Tools analysiert, hab den Gap sozusagen weiter erklärt und hab dann gesagt:

00:06:43,940 --> 00:06:45,040 [speaker_2]
Mhm.

00:06:45,100 --> 00:07:09,460 
Genau, und das wäre sozusagen eine, ein Sample, ein Sample Environment. Du hast halt,
könntest ihn rein theoretisch auf einen Server mit Applications draufschmeißen,

00:07:09,490 --> 00:07:09,490 [speaker_2]
Mhm.

00:07:10,280 --> 00:08:26,880 
Genau, es gibt sozusagen jetzt dann zwei große Konfigurationstypen, und zwar einen Secret Store.
Der ist sozusagen: Wie komme ich zum Vault hin? Wie kann ich mich authentifizieren mit dem Vault?

00:08:27,600 --> 00:08:27,920 [speaker_2]
Mhm.

00:08:28,339 --> 00:16:25,736 
Und das würde der ESO gerade noch nicht machen. Das heißt,
du hättest eine broken Configuration Datei. Ich pro--

00:16:27,136 --> 00:16:35,916 
Okay, äh, Fragen per se habe ich nicht. Ich fand's, ich find's irgendwie ganz interessant. Äh,
kannst du's mir noch in den Chat hineingeben?

00:16:35,976 --> 00:16:36,556 
Ja, ja, natürlich.

00:16:36,566 --> 00:16:37,296 
Kann ich das abtippen.

00:16:38,436 --> 00:16:49,676 
[lacht] Ja, bitte tipp das ab. Aber davor, warte. Warte, davor: Wenn dich, wenn's dich interessiert,
hier ist er sozusagen gemonitort. Also hier sieht man,

00:16:50,036 --> 00:16:50,096 
Mhm.

00:16:50,156 --> 00:17:01,896 
Er hat hundert Secrets zu managen at the moment und wir sehen,
ohne Last braucht er null Komma drei CPU, mit Last null Komma vier, fünf. Also du merkst,

00:17:02,376 --> 00:17:02,556 
Mhm.

00:17:02,576 --> 00:17:04,406 
Und ich find, das ist eigentlich relativ effizient-

00:17:04,536 --> 00:17:04,546 
Ja.

00:17:04,546 --> 00:17:14,296 
-dafür, dass der hundert Files macht. Und was braucht er? Er allocated sich vierzig MB.
Brauchen tut er die ganze Zeit, selbst bei den Secret Rerolls, irgendwas um die fünf, sechs MB.

00:17:14,976 --> 00:17:26,335 
Mhm. Äh, hast du geschaut gehabt jemals, ob die, äh, ob die,
der Memory Demand auch zum Beispiel bei einem Vielfachen davon,

00:17:26,516 --> 00:17:34,356 
Nein, habe ich-- Nein, habe ich noch nicht, aber die Sache ist per se, du hast,
also kein Host hat eine Million Secrets. Das ist ein-

00:17:35,576 --> 00:17:53,256 
Ah. Also das, das große Problem, das ich gesehen hatte oftmals, ist: Nein,
du brauchst keine eine Million Secrets. Niemand braucht eine Million Secrets. Das Problem ist, äh,

00:17:54,836 --> 00:18:06,895 
Ich verwend, äh, intern mit Go eins sechsundzwanzig das Secrets Punkt do, was sozusagen, sobald der,
sobald er merkt, dass er das wegräumen kann, alle Register aus der CPU rauslöscht von dem-

00:18:06,956 --> 00:18:27,416 
Absolut, absolut. Aber stell dir das Szenario vor, jemand kommt und sagt,
er createt ein neues Secret und benutzt dieses neue Secret in seinem Service,

00:18:27,476 --> 00:18:29,656 
Und das ist halt ein Prototyp. Das ist ja bei einer Master-

00:18:29,816 --> 00:18:30,776 
Ja, ja, absolut.

00:18:30,796 --> 00:18:33,636 
-baust du ja kein Production-Grade System, sondern baust einmal was und schaust.

00:18:33,696 --> 00:18:43,918 
Ja, ja. Die Idee, die Idee finde ich ganz cool. Die Idee finde ich cool.Genau. Äh, das heißt,
dein Prototyp funktioniert also jetzt, hast du gesagt, mit deinem normalen Linux-Host,

00:18:43,948 --> 00:18:58,348 
Ja, genau. Das Kubernetes-Dings ist noch immer der External Secret Operator.
Da hat sich nichts geändert. Warum sollte ich den noch mal machen? Sondern du hast einfach,

00:18:59,088 --> 00:19:06,668 
Das heißt, du hast eigentlich deinen Secret Agent, der managed nur für beide Maschinen,
weil die ja nicht kompatibel ist mit der Linuxmaschine.

00:19:07,588 --> 00:19:10,248 
Nein, nein. Ich habe nur einen Agent auf Linux geschrieben.

00:19:11,548 --> 00:19:17,828 
Das, das meinte ich. Also dein Linux Agent handelt eben das,
dass er ja normalerweise nicht wirklich gut kompatibel mit dem-

00:19:18,348 --> 00:19:49,128 
Genau, es ist sozusagen conceptually aligned, um halt sozusagen den, das Ganze abzubilden.
Ich habe dir jetzt auch noch das Repository geschickt.

00:19:49,348 --> 00:19:49,728 
Mhm.

00:19:50,628 --> 00:19:59,148 
Genau,
das ist sozusagen einfach nur hier dokumentiert und sonst ist das Ding halt Open Source und das

00:19:59,918 --> 00:20:00,108 
Mhm.

00:20:00,268 --> 00:20:15,808 
Genau, jetzt wünsche ich dir viel Spaß bei den Fragen von den Forms.
Und dann kommen wir noch zu offenen Fragen,

00:20:15,848 --> 00:21:29,608 
Das ist ... [räuspert sich] Äh, das ist noch eine Frage. Du sagst, äh,
das ist vor allem bei dieser Frage jetzt:

00:21:30,188 --> 00:21:55,068 
Doch, die Secrets per se liegen dann als plain, also liegen dann als plain text,
als Datei am Host herum, vom Secret User geowned, von der Gruppe lesbar oder vom Secret User,

00:21:55,688 --> 00:22:05,208 
Ja, ja, ja. Nein, für mich war es zuerst jetzt dieses, dass es fixe, fertige Secrets sind,
die einfach nur da sind und für immer dort sind. Aber das heißt durch das Rotation,

00:22:05,468 --> 00:22:17,468 
Genau. Dadurch, dadurch, dass die halt aufgelöst werden, diese Content-Sachen, hast du halt so lang,
wie die time to leave ist, ist dieses File halt an der Stelle, wo du konfiguriert hast am Host.

00:22:17,848 --> 00:22:26,968 
Hast du dir überlegt, ob es da irgendeinen Nutzen bringen würde,
diese Dateien zumindest für die Zeit, in der sie da sind, äh, zu verschlüsseln?

00:22:27,348 --> 00:22:31,768 
Die Sache ist, dann musst du der Applikation ja wieder den Encryption Key irgendwie geben.

00:22:32,348 --> 00:22:32,588 
Ja.

00:22:33,388 --> 00:22:33,698 
Und dann-

00:22:33,808 --> 00:22:39,028 
Dann müsstest du ja ein Handshake-Verfahren am Anfang haben, das wiederum mit ... Ja, klar.

00:22:39,088 --> 00:22:44,128 
Das macht halt das Ganze-- Also wie gesagt,
du hast halt immer diesen Security-Usability-Trade-off am Ende des Tages.

00:22:44,168 --> 00:22:45,348 
Ja, ja, ja. Nein, ist schon.

00:22:45,948 --> 00:22:52,348 
Und wobei man halt-- Ich sag halt auch, sobald jemand Host, also Root-Zugriff hat,
kann er einen Memory Dump von der Applica-, von der Applikation auch machen.

00:22:52,368 --> 00:22:58,508 
Sowieso. Und nicht nur das,
sondern dann kannst du auch ganz ehrlich eher auch die Keys selbst dir holen.

00:22:58,588 --> 00:22:58,648 
Ja.

00:22:59,388 --> 00:23:00,708 
Daher, ja.

00:23:00,728 --> 00:23:04,948 
Demnach, wie gesagt, wenn dus ganz secure haben willst,
musst du es halt mit Managed Identities lösen.

00:23:05,008 --> 00:23:05,728 
Genau. Ja.

00:23:06,228 --> 00:23:09,048 
Aber das ist halt einfach ein anderer Use Case.

00:24:32,008 --> 00:24:32,308 
Passt.

00:24:33,148 --> 00:24:44,484 
Wunderbar. Gut, dann kommen wir zu den Fragen-Hast du bei den Fragen irgendeine Frage gehabt,
wo du gerne was aufgrund deiner Antwort sagen möchtest oder nicht?

00:24:44,624 --> 00:24:53,624 
Mhm, ich, also ich, ich bin eigentlich zufrieden mit dem, wie du's gemacht hast.
Daher hab ich jetzt auch nicht wirklich, ne. Also die, die Frage, die ich gehabt hatte,

00:24:53,764 --> 00:25:03,284 
Wunderbar. Was siehst du sozusagen als Hauptvorteile von so einem Linux Secret Agent also?
Oder besser gesagt von dem Ligu-- Linux Secret Agent, den ich gebaut hab?

00:25:04,604 --> 00:25:42,344 
Einer der Hauptvorteile, die ich sehe, wär die Flexibility. Weil eines der großen Probleme,
die man immer wieder hat, ist, äh, das Login von gewissen Environments. Vor allem,

00:25:42,843 --> 00:25:53,444 
Genau, was du, was sozusagen auch cool ist: Du kannst Brownfield und gekaufte Systeme anbinden,
anbinden an deine Secret-Management-Infrastruktur, solang die halt Files konsumieren können.

00:25:53,504 --> 00:25:54,084 
Stimmt, ja.

00:25:54,124 --> 00:25:58,864 
Also solang es von Files ist, kannst du sozusagen alles eigentlich abbilden und ja, genau.

00:25:58,984 --> 00:25:59,264 
Mhm.

00:25:59,784 --> 00:26:05,964 
Den Trade-off, den hast du eh schon gea, beantwortet, dass das für dich ein okayer Trade-off ist,
dass die Da, dass die Datei halt am Host existiert.

00:26:06,124 --> 00:26:17,684 
Ja, es ist, äh, ich sag's so: Ich hab's jetzt als, als, äh, das Zweitbeste auch angekreuzt gehabt.
Für mich war das so: Ja, es ist ein Trade-off,

00:26:17,764 --> 00:26:18,143 
Ja.

00:26:18,154 --> 00:26:38,324 
Aber es ist ein absolut acceptable Trade-off, wo ich sage: Gut. [seufzt] Es, es,
es gibt halt immer so ein bisschen so 'n, so 'n Punkt auch, das ist genau diese Meinung,

00:26:39,344 --> 00:26:52,144 
Genau. Was ich sozusagen relativ gut an dem System finde: Selbst wenn ein Host kompromittiert ist,
sobald du's mitbekommst, kannst du halt per se die, die Reg--

00:26:53,684 --> 00:26:53,964 
Mhm.

00:26:54,464 --> 00:27:24,824 
Und per se sind dann nur alle Secrets bis zu genau dem Moment retrievable.
Dann kannst du sozusagen alle rollieren, den neue VM über Terraform, was auch immer,

00:27:24,924 --> 00:27:44,184 
Ja. Wichtig bei dem Ganzen ist aber dann für mich, ähm,
das war ja auch ein paar Mal ist das vorgekommen: So wie du's jetzt dargestellt hast,

00:27:44,844 --> 00:27:53,184 
Genau. Ja, das ist, ja, da hast du recht. Es ist kein End-to-End-System, was alles abnimmt,
sondern ein kleiner Teil im Secret-Management-Kosmos.

00:27:53,684 --> 00:27:53,844 
Mhm.

00:27:53,964 --> 00:28:05,304 
Genau. Gut. Das be-- Gibt's irgendwelche Bereiche als nur wirklich Application Secret Management,
wo du glauben würdest, dass dieses Konzept, dieser Agent einsetzbar wäre?

00:28:06,964 --> 00:28:25,504 
Es kommt ganz darauf an, was, äh-- Du hast ja gesagt,
du hast den Secret Agent und der gibt dir die Secrets mehr oder weniger. Ähm, wie,

00:28:26,004 --> 00:28:32,184 
Du bekommst sie sozusagen genau, wie du's hier schreibst, kriegst du's zurück. Das heißt,
du schreibst hier deine Datei.

00:28:32,824 --> 00:28:33,104 
Mhm.

00:28:33,724 --> 00:29:01,724 
Warte, sonst öffne ich kurz einfach meine ID, dann können wir uns das noch besser anschauen.
Also du schreibst halt einfach den, das Template in der Datei und den genauen Pfad. Das heißt,

00:29:02,184 --> 00:29:09,623 
Ja, ja, das heißt, du kriegst-- Ich sehe es gerade. Das heißt,
du kriegst mehr oder weniger die Datei als solche kriegst du einfach zurück.

00:29:09,964 --> 00:29:10,164 
Genau.

00:29:10,184 --> 00:29:32,724 
Das heißt, ich kann auch beispielsweise – das ist zum Beispiel eine Sache, wo,
was mich auch interessiert, ob ich direkt mit Docker, wenn ich beispielsweise jetzt, sagen wir mal,

00:29:32,784 --> 00:29:32,884 
Ja.

00:29:33,164 --> 00:29:43,264 
Könnte ich dann die Secrets, wie zum Beispiel den, äh, die Datenbank,
den Datenbankpasswort und so weiter im Compose-File direkt als Environment Variable hineinhauen?

00:29:43,964 --> 00:29:47,544 
Die Sache ist, es gibt ja die, es gibt ja im Punkt Compose auch diese Punkt Endfiles.

00:29:47,944 --> 00:29:48,304 
Genau.

00:29:48,804 --> 00:29:52,764 
Das heißt, du könntest dir sozusagen über diesen Agent dieses Punkt Endfile generieren lassen.

00:29:52,844 --> 00:29:53,884 
Okay, passt. Mhm.

00:29:53,924 --> 00:29:55,814 
Und dann halt das Starten sozusagen dieses-

00:29:55,824 --> 00:29:57,414 
Genau, das Restarten dann, wenn-

00:29:57,444 --> 00:30:00,584 
Ja, das Restarten, wenn, musstes, müsstest du halt selber lösen.

00:30:00,744 --> 00:30:01,044 
Ja, ja.

00:30:01,224 --> 00:30:14,364 
Das ist auch absichtliche Con-consideration. Man könnte natürlich überlegen, ob man einen,
einen Hook Schrägstrich Signal Endpoint Konfiguration haben sollte. Und ich hab mir gedacht,

00:30:14,384 --> 00:30:38,824 
Genau, das hab ich auch beim Punkt. Ja, ich kann ja, kann ja das Environment File und jedes Mal,
wenn sich irgendwas ändert, dann restarte ich einfach. Aber das passt, das ist cool. Ja, nee, das,

00:30:39,184 --> 00:30:39,324 
Ja.

00:30:39,904 --> 00:30:52,464 
Und du willst jetzt auf diesem Team dein kleines Mini Service hosten und das ist schon
containerized. Und was man dann halt oft sieht, ist dann,

00:30:52,684 --> 00:30:53,064 
Ja.

00:30:53,624 --> 00:30:59,104 
Und da wär's dann halt interessant,
ob man sich trotzdem anbinden kann und das Secret in-inkludieren könnte.

00:30:59,624 --> 00:31:08,444 
Ja, kann man zum Glück. Also insofern, dass du sozusagen, äh, davor weißt, wo das File sein wird,
geht das.

00:31:09,404 --> 00:31:11,084 
Soll das sitzen?

00:31:11,124 --> 00:31:19,144 
Ja, genau. Einfach hier können wir uns-- Also hier für sich ist ein Audit Log. Das heißt,
hier sieht man die File Rewrites, die er macht.

00:31:19,224 --> 00:31:20,243 
Mhm.

00:31:21,304 --> 00:31:41,664 
Und wenn wir jetzt zum Beispiel-- So, hier sagt uns E1, das heißt, hier wurde einfach konfiguriert.
Über diese Datei hab ich ihm halt gesagt: „Generier mir das und das soll bitte so ausschauen“,

00:31:42,084 --> 00:31:42,174 
Mhm.

00:31:42,264 --> 00:32:07,684 
Hier kann ich halt-- Hier sage ich ihm halt in dem Fall, weil das jetzt ein,
ein HashiCorp Vault ist und kein Azure Key Vault, habe ich halt hier die hierarchischen Pfade,

00:32:08,084 --> 00:32:08,994 
Ja, ja, ich verstehe.

00:32:08,994 --> 00:32:11,364 
Ich brauche nicht beweisen, dass die Go-Templating-Language funktioniert.

00:32:12,204 --> 00:32:13,064 
Wirklich nicht.

00:32:13,104 --> 00:32:18,104 
Haben andere bewiesen. Genau. Und so funktioniert das halt und so, das macht er halt. Genau.

00:32:18,124 --> 00:32:21,764 
Ja, ist cool. Passt. Danke. Schaut gut aus.

00:32:23,984 --> 00:32:39,204 
Genau. Was hältst-- Äh, das ist jetzt nicht mehr so wichtig, dö-dö. Siehst du eine--
Was ist der größte Benefit Schrägstrich was für einen Benefit siehst du,

00:32:43,424 --> 00:33:38,224 
Einer der größten Benefits? Na ja,
von conceptionally aligned ist es halt eben vor allem dieses Du bist nicht-- Ach, ich weiß nicht,

00:33:39,124 --> 00:33:39,134 
Mhm.

00:33:39,284 --> 00:34:00,284 
Ich finde, das ist eines der größten Benefits, ist auch die Verständlichkeit. Sprich,
wenn man die Basics versteht, wie ein Linux-System funktioniert, wie Filesystems funktionieren, wie,

00:34:00,924 --> 00:34:28,773 
Genau. Ja, was ich persönlich so cool finde, ist, du kannst Developern die Möglichkeit geben,
ihre Konfiguration mit Secrets zu managen, ohne dass die jemals ein Secret sehen müssen.

00:34:29,384 --> 00:34:29,584 
Ja.

00:34:29,604 --> 00:34:38,604 
Das legt er an, gibt ihm den Pfad oder editiert dann selber den Pfad in die Config rein,
in den Merge Request, je nachdem, wie halt man den Workflow aufbauen will und gut ist. Genau.

00:34:38,704 --> 00:34:39,684 
Mhm.

00:34:40,804 --> 00:34:44,714 
Dö-dö-dö. Also ja. Also ich sehe einen großen Vorteil in dieser, in die, also-

00:34:44,784 --> 00:35:05,604 
Nein, absolut. I-I-Ich sag gleich, ich bin, ich bin ein großer Fan von keeping it simple,
sehr die Standard-Tools von Linux verwenden.

00:35:05,884 --> 00:35:10,404 
Man verwendet Just aber, ja. Oder Justify heißt-- Justify heißt das bessere Make, aber ja.

00:35:11,324 --> 00:35:17,504 
Ja, aber, äh, erstens, wenn man schon so lange dabei ist, dann weiß,
dann hat man sehr viele Makefiles schon irgendwo herumliegen.

00:35:18,004 --> 00:35:18,014 
Ja.

00:35:18,024 --> 00:35:23,104 
Und eines der Hauptdinge ist halt immer dieses: Make ist fast immer bereits installiert.

00:35:23,794 --> 00:35:23,994 
Ja.

00:35:23,994 --> 00:35:32,364 
Most Dev-Tools sind immer installiert und da ist Make immer dabei. Und das ist eines der Sachen:
Du hast nicht immer Admin-Rechte, um irgendwelche Tools noch extra zu installieren.

00:35:32,404 --> 00:35:32,644 
Mhm.

00:35:32,904 --> 00:35:43,984 
Und in dem Fall ist es halt genau so ein Ding,
wo man halt nicht irgendwelche kommerziellen noch Lösungen haben musst und keine Ahnung was,

00:35:44,484 --> 00:36:05,284 
Verstehe. Wunderbar, vielen Dank. So,
gibt's sozusagen Improvements oder Feature Requests oder irgendwas,

00:36:06,664 --> 00:36:15,784 
Äh, die Vereinfachung des Deployments-- Oder ne, ne, die Deployments find,
fand ich eh ganz cool mit den, mit den, mit den, ähm-

00:36:16,404 --> 00:36:17,044 
Ansible.

00:36:17,104 --> 00:37:16,260 
Ansible-Files, äh, mit der Playbooks. Äh, ich glaube, das was,
wo man wirklich noch vielleicht arbeiten müsste, dass man das production ready hat, ist,

00:37:17,420 --> 00:37:25,730 
Verstehe. Ich verstehe, ziehst du meinst, ich fasse es kurz zusammen: Du meinst sozusagen ganz,
aber das, das kann man halt nicht wirklich machen,

00:37:25,760 --> 00:37:48,920 
Natürlich, natürlich. Aber man kann zum Beispiel die,
die Standard Cases sollte man schon zum Beispiel abdecken. Beispielsweise wie das,

00:37:49,460 --> 00:37:56,140 
Ja, wie gesagt, da könnte man konzeptionell schon Input geben,
aber per se musst du's dann halt pro Secret Store und pro Environment selber lösen.

00:37:56,720 --> 00:37:57,900 
Ja, ja.

00:37:57,960 --> 00:37:58,500 
Ja.

00:37:58,560 --> 00:38:02,560 
Je, je einfacher das passiert, desto eher kann man von Production Readiness reden.

00:38:02,600 --> 00:38:17,360 
Verstehe. Danke für den Insight. Genau. Könnte, also wahrscheinlich wär auch das,
also gibt's in deiner Organisation, in der du grad arbeitest, irgendwas,

00:38:19,220 --> 00:38:49,540 
Äh, vor allem hier an der Uni nicht explizit. Wir haben ja oftmals,
arbeiten wir mit verschiedensten Linux VMs et cetera. Würde jetzt nicht explizit was blockieren.

00:38:50,520 --> 00:38:59,020 
Verstehe natürlich und darum mache ich auch Interviews, weil DevSecOps,
wurde auch in einigen Papers zitiert: „It's not about tools, it's about people."

00:38:59,700 --> 00:38:59,760 
Ja.

00:39:00,040 --> 00:39:03,960 
Du kannst das allerbeste Tool bauen, aber wenn das zu komplex ist oder das keiner verwenden möchte.

00:39:04,020 --> 00:39:42,840 
Ja. Das war, das war deshalb auch mein Kommentar mit dem keeping it simple ist toll. Ähm, das heißt,
du kannst auch vieles von dem automatisieren mit Skripten, Playbooks, was auch immer, ähm,

00:39:43,580 --> 00:39:50,380 
Das ist ja eh möglich. Du hast ja sozusagen,
du spielst ja pro Installation spielst du ja deine eigenen Secret Stores drauf, so wie wir-

00:39:50,540 --> 00:39:55,620 
Nein, nein, ich meinte, ich meinte zum Beispiel, äh, äh, äh, ähm, Catastrophic Recovery.

00:39:55,720 --> 00:39:56,720 
Ah, das meinst du.

00:39:57,000 --> 00:40:40,400 
Also beispielsweise, wenn du so eine Situation hast, dann hast du ein Skript,
eine Vorlage und sagst: „Okay, das muss passieren. In dem Skript passiert das",

00:40:40,980 --> 00:40:45,920 
Verstehe. Na gut, ich glaube, du kannst einmal das Recording stoppen. Jetzt können wir noch so fünf,
zehn Minuten kurz ein bisschen reden und dann-
